#!/usr/bin/env python3
"""
Inlines Tyk's OpenAPI 3.0 document schema (tyk/apidef/oas/schema/3.0.json)
into a Gateway or Dashboard swagger (see tyk-docs.yml, the "gateway" and
"dashboard" jobs).

The source swaggers reference that schema by an external raw.githubusercontent
URL. Mintlify cannot resolve external $refs, and client generators cannot
resolve it offline. Replacing it with an empty object keeps Mintlify happy but
makes generated clients drop openapi/info/paths from Tyk OAS API models.

3.0.json is JSON Schema draft-04, so its definitions are converted to OAS 3.0
schema objects and added as components.schemas.OAS3*:
  - $schema, id and definitions are dropped
  - patternProperties becomes additionalProperties
  - the ExampleXORExamples / SchemaXORContent helpers are dropped (they only
    constrain, they carry no data shape and break some generators)

Only the $ref lines change and the new schemas are inserted as a block under
components.schemas, so the rest of the swagger stays byte-identical.
"""
import argparse
import json
import re
import sys

import yaml

EXTERNAL_REF = re.compile(
    r"""['"]?https://raw\.githubusercontent\.com/TykTechnologies/tyk/refs/(?:heads|tags)/[^/'"\s]+/apidef/oas/schema/3\.0\.json['"]?"""
)
PREFIX = "OAS3"
ROOT = PREFIX + "Document"
HELPERS = ("ExampleXORExamples", "SchemaXORContent")


def is_helper_ref(node):
    return isinstance(node, dict) and str(node.get("$ref", "")).endswith(HELPERS)


def convert(node):
    if isinstance(node, list):
        return [convert(v) for v in node]
    if not isinstance(node, dict):
        return node

    out = {k: convert(v) for k, v in node.items() if k not in ("$schema", "id", "definitions")}

    ref = out.get("$ref")
    if isinstance(ref, str) and ref.startswith("#/definitions/"):
        return {"$ref": "#/components/schemas/" + PREFIX + ref.rsplit("/", 1)[-1]}

    if isinstance(out.get("allOf"), list):
        out["allOf"] = [v for v in out["allOf"] if not is_helper_ref(v)]
        if not out["allOf"]:
            del out["allOf"]

    pattern_props = out.pop("patternProperties", None)
    if pattern_props:
        named = [v for k, v in pattern_props.items() if k != "^x-"]
        if named:
            out["additionalProperties"] = named[0]
        elif out.get("additionalProperties") is False:
            out["additionalProperties"] = True  # allow x- extensions

    return out


def build_schemas(schema):
    schemas = {
        PREFIX + name: convert(definition)
        for name, definition in schema["definitions"].items()
        if not name.endswith(HELPERS)
    }
    schemas[ROOT] = convert(schema)
    return schemas


def main():
    parser = argparse.ArgumentParser(description="Inline tyk/apidef/oas/schema/3.0.json into a swagger file")
    parser.add_argument("--input-file", required=True, help="Swagger YAML to patch in place")
    parser.add_argument("--schema-file", required=True, help="Path to tyk/apidef/oas/schema/3.0.json")
    args = parser.parse_args()

    with open(args.input_file, "r", encoding="utf-8") as f:
        content = f.read()

    content, count = EXTERNAL_REF.subn(f"'#/components/schemas/{ROOT}'", content)
    if count == 0:
        print(f"ℹ️ No external 3.0.json $ref in {args.input_file}, nothing to inline")
        return

    existing = yaml.safe_load(content)["components"]["schemas"]
    collisions = sorted(k for k in existing if k.startswith(PREFIX))
    if collisions:
        print(f"❌ {args.input_file} already defines {collisions}, update PREFIX in this script")
        sys.exit(1)

    with open(args.schema_file, "r", encoding="utf-8") as f:
        schemas = build_schemas(json.load(f))

    block = yaml.safe_dump(schemas, sort_keys=False, allow_unicode=True, width=1000)
    block = "".join("    " + line if line.strip() else line for line in block.splitlines(keepends=True))

    anchor = re.search(r"^components:\n(?:.*\n)*?  schemas:[ \t]*\n", content, re.MULTILINE)
    if not anchor:
        print(f"❌ components.schemas not found in {args.input_file}")
        sys.exit(1)
    content = content[: anchor.end()] + block + content[anchor.end() :]

    patched = yaml.safe_load(content)
    all_schemas = patched["components"]["schemas"]
    missing = sorted(
        {r.rsplit("/", 1)[-1] for r in re.findall(r"#/components/schemas/(" + PREFIX + r"\w+)", content)} - set(all_schemas)
    )
    if missing:
        print(f"❌ Dangling refs after inlining: {missing}")
        sys.exit(1)

    with open(args.input_file, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"✅ Replaced {count} external 3.0.json refs and added {len(schemas)} {PREFIX}* schemas in {args.input_file}")


if __name__ == "__main__":
    main()
