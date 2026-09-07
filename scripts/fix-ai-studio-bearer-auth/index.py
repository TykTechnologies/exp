#!/usr/bin/env python3
"""
Patches the AI Studio swagger's BearerAuth security scheme after the
Swagger 2.0 -> OpenAPI 3.0 conversion (see tyk-docs.yml, the "ai-studio" job).

Swagger 2.0 has no native "http bearer" scheme type, so ai-studio's
swaggo-generated spec models bearer auth as:

    BearerAuth:
      type: apiKey
      name: Authorization
      in: header

which the converter carries forward unchanged. type:apiKey tells any
spec-consuming tool (Mintlify's API playground, codegen, Postman) to send
the header value as-is, with no prefix. The AI Studio server actually
requires "Authorization: Bearer <key>", so those tools send a request that
gets a 401.

This patches the scheme to the correct OpenAPI 3.x form:

    BearerAuth:
      type: http
      scheme: bearer

which every consumer of the spec knows how to prefix correctly. Remove this
script (and its call in tyk-docs.yml) once ai-studio emits this correctly
in its own generated swagger.
"""
import argparse
import sys

OLD_BLOCK = (
    "  securitySchemes:\n"
    "    BearerAuth:\n"
    "      type: apiKey\n"
    "      name: Authorization\n"
    "      in: header\n"
)

NEW_BLOCK = (
    "  securitySchemes:\n"
    "    BearerAuth:\n"
    "      type: http\n"
    "      scheme: bearer\n"
)


def main():
    parser = argparse.ArgumentParser(
        description="Fix AI Studio swagger's BearerAuth scheme to type:http/scheme:bearer"
    )
    parser.add_argument("--input-file", required=True, help="Path to the converted OpenAPI 3.x YAML file")
    args = parser.parse_args()

    with open(args.input_file, "r", encoding="utf-8") as f:
        content = f.read()

    if OLD_BLOCK not in content:
        print("❌ Expected BearerAuth type:apiKey block not found - ai-studio's security scheme may have changed, update this script:")
        print(content)
        sys.exit(1)

    content = content.replace(OLD_BLOCK, NEW_BLOCK, 1)

    with open(args.input_file, "w", encoding="utf-8") as f:
        f.write(content)

    print(f"✅ Patched BearerAuth to type:http/scheme:bearer in {args.input_file}")


if __name__ == "__main__":
    main()
