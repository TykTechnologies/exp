#!/usr/bin/env python3
"""Generate a dot-notation Mintlify SNIPPET documenting all Tyk Operator CRDs.

Reads CRD YAMLs (config/crd/bases/*.yaml) and emits flattened <ParamField> rows
for every field under each CRD's `spec`. Source of truth: the OpenAPI v3 schema
embedded in each CRD, so we get description + type + enum + default + required.

Output is a frontmatter-less snippet (snippets/operator-crd-reference.mdx) that is
imported by product-stack/tyk-operator/crd-reference.mdx, which owns the page
frontmatter and intro. The sync workflow overwrites only this snippet.

Per-CRD layout (matches the approved dot-notation sample):
  ## <Kind>
  ### Top-Level Fields        -> the CRD's direct scalar/leaf spec fields
  ### <objectField>           -> one H3 per top-level object field, with its subtree

Ordering at every level: required fields first, then alphabetical by name.
"""
import sys, os, glob, yaml

# Preferred CRD ordering on the page (Tyk OAS first per docs style); rest appended alphabetically.
ORDER = [
    "TykOasApiDefinition", "ApiDefinition", "TykStreamsApiDefinition",
    "TykMcpProxyDefinition", "SecurityPolicy", "APIDescription",
    "OperatorContext", "PortalAPICatalogue", "PortalConfig",
    "SubGraph", "SuperGraph",
]

def esc(text):
    """Make a CRD description safe to drop into MDX (JSX) content."""
    if not text:
        return ""
    t = " ".join(str(text).split())          # collapse newlines/whitespace
    t = t.replace("&", "&amp;")
    t = t.replace("<", "&lt;").replace(">", "&gt;")
    t = t.replace("{", "&#123;").replace("}", "&#125;")
    return t

def attr_esc(v):
    return str(v).replace('"', "&quot;")

def sort_key(name, required_set):
    # required first (False sorts before True), then alphabetical
    return (name not in required_set, name.lower())

def is_objectish(sch):
    """A field that should get its own H3 / be recursed into (has a real sub-schema)."""
    t = sch.get("type")
    if t == "array":
        return is_objectish(sch.get("items", {}) or {})
    if t == "object" or (t is None and ("properties" in sch or "additionalProperties" in sch)):
        if sch.get("properties"):
            return True
        if isinstance(sch.get("additionalProperties"), dict):
            return True
        return False  # free-form object -> single leaf row
    return False

class Gen:
    def __init__(self):
        self.lines = []
        self.count = 0

    def param(self, path, typ, required, desc, default=None):
        self.count += 1
        head = f'<ParamField body="{path}" type="{typ}"'
        if required:
            head += " required"
        if default is not None:
            head += f' default="{attr_esc(default)}"'
        d = esc(desc)
        if d:
            self.lines.append(head + ">")
            self.lines.append(f"  {d}")
            self.lines.append("</ParamField>")
        else:
            self.lines.append(head + " />")
        self.lines.append("")

    def field(self, path, sch, required):
        """Entry point for one named field; unwraps arrays then delegates to node()."""
        if sch.get("type") == "array":
            items = sch.get("items", {}) or {}
            self.node(path + "[]", items, required, desc_override=sch.get("description"))
        else:
            self.node(path, sch, required)

    def node(self, path, sch, required, desc_override=None):
        t = sch.get("type")
        desc = desc_override if desc_override is not None else sch.get("description")
        is_obj = t == "object" or (t is None and ("properties" in sch or "additionalProperties" in sch))

        if is_obj:
            ap = sch.get("additionalProperties")
            if isinstance(ap, dict):  # map[string]X
                self.param(path, "object", required, desc or "Map of string keys to values.")
                self.field(path + ".{key}", ap, False)
                return
            props = sch.get("properties")
            if props:
                self.param(path, "object", required, desc)
                req = set(sch.get("required", []))
                for name in sorted(props, key=lambda n: sort_key(n, req)):
                    self.field(path + "." + name, props[name], name in req)
                return
            note = "Free-form object with arbitrary keys."
            self.param(path, "object", required, (desc + " " if desc else "") + note)
            return

        # leaf
        typ = t or "object"
        if "enum" in sch:
            vals = ", ".join("`%s`" % v for v in sch["enum"])
            desc = (desc + " " if desc else "") + "**Allowed values:** " + vals + "."
        self.param(path, typ, required, desc, default=sch.get("default"))

def storage_version(versions):
    for v in versions:
        if v.get("storage"):
            return v
    return versions[-1]

def main(yaml_dir, out_path):
    crds = {}
    for f in sorted(glob.glob(os.path.join(yaml_dir, "*.yaml"))):
        doc = yaml.safe_load(open(f))
        kind = doc["spec"]["names"]["kind"]
        ver = storage_version(doc["spec"]["versions"])
        schema = ver["schema"]["openAPIV3Schema"]
        spec = schema.get("properties", {}).get("spec", {})
        crds[kind] = {
            "schema_desc": schema.get("description"),
            "spec_desc": spec.get("description"),
            "props": spec.get("properties", {}) or {},
            "required": set(spec.get("required", [])),
        }

    ordered = [k for k in ORDER if k in crds] + sorted(k for k in crds if k not in ORDER)

    # This file is a Mintlify SNIPPET (no frontmatter): the generated body only.
    # It is imported by product-stack/tyk-operator/crd-reference.mdx, which owns
    # the frontmatter and intro. The sync workflow overwrites only this snippet.
    out = []
    out.append("{/* DO NOT EDIT. Auto-generated from the Tyk Operator CRD schemas (config/crd/bases) by the exp tyk-docs sync workflow. To change a field description, edit the Go doc-comments in tyk-operator-internal. */}")
    out.append("")

    total = 0
    for kind in ordered:
        c = crds[kind]
        out.append(f"## {kind}")
        out.append("")
        intro = c["spec_desc"] or c["schema_desc"]
        if intro:
            out.append(esc(intro))
            out.append("")
        if not c["props"]:
            out.append("_This resource has no user-configurable `spec` fields._")
            out.append("")
            continue

        leaves, objects = [], []
        for name, sch in c["props"].items():
            req = name in c["required"]
            (objects if is_objectish(sch) else leaves).append((name, sch, req))
        leaves.sort(key=lambda x: (not x[2], x[0].lower()))
        objects.sort(key=lambda x: (not x[2], x[0].lower()))

        if leaves:
            out.append("### Top-Level Fields")
            out.append("")
            g = Gen()
            for name, sch, req in leaves:
                g.field(name, sch, req)
            out.extend(g.lines)
            total += g.count

        for name, sch, req in objects:
            out.append(f"### {name}")
            out.append("")
            g = Gen()
            g.field(name, sch, req)
            out.extend(g.lines)
            total += g.count

    with open(out_path, "w") as fh:
        fh.write("\n".join(out).rstrip() + "\n")

    print(f"CRDs: {len(ordered)}")
    print(f"Total ParamField rows: {total}")
    print("Order: " + ", ".join(ordered))

if __name__ == "__main__":
    main(sys.argv[1], sys.argv[2])
