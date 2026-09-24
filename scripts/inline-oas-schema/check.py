#!/usr/bin/env python3
"""
Fails if a Tyk OAS API body schema in a swagger lost its OpenAPI fields.

Every allOf that includes the Tyk vendor extension must have a member that
defines `openapi`. An empty placeholder (type: object, additionalProperties:
true) passes spec validation but makes generated clients drop
openapi/info/paths, so catch it here.
"""
import argparse
import sys

import yaml

# Older swaggers use XTykApiGateway directly instead of TykVendorExtension.
VENDOR_REFS = {"#/components/schemas/" + n for n in ("TykVendorExtension", "XTykApiGateway", "XTykAPIGateway")}


def walk(node, path=""):
    if isinstance(node, dict):
        yield path, node
        for k, v in node.items():
            yield from walk(v, f"{path}/{k}")
    elif isinstance(node, list):
        for i, v in enumerate(node):
            yield from walk(v, f"{path}[{i}]")


def resolve(schemas, ref):
    if not ref.startswith("#/components/schemas/"):
        return {}
    return schemas.get(ref.rsplit("/", 1)[-1], {})


def main():
    parser = argparse.ArgumentParser(description="Check Tyk OAS API body schemas keep their OpenAPI fields")
    parser.add_argument("--input-file", required=True)
    args = parser.parse_args()

    with open(args.input_file, "r", encoding="utf-8") as f:
        spec = yaml.safe_load(f)
    schemas = spec["components"]["schemas"]

    checked, failures = 0, []
    for path, node in walk(spec):
        all_of = node.get("allOf")
        if not isinstance(all_of, list) or not any(isinstance(m, dict) and m.get("$ref") in VENDOR_REFS for m in all_of):
            continue
        checked += 1
        refs = [m.get("$ref", "") for m in all_of if isinstance(m, dict)]
        if not any("openapi" in resolve(schemas, r).get("properties", {}) for r in refs):
            failures.append(f"{path}: none of {refs} defines `openapi`")

    if checked == 0:
        print(f"❌ No Tyk OAS API schemas found in {args.input_file}, update VENDOR_REFS in this script")
        sys.exit(1)
    if failures:
        print(f"❌ {len(failures)} Tyk OAS API schemas lost their OpenAPI fields in {args.input_file}:")
        print("\n".join(failures))
        sys.exit(1)
    print(f"✅ {checked} Tyk OAS API schemas keep their OpenAPI fields in {args.input_file}")


if __name__ == "__main__":
    main()
