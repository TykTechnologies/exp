#!/usr/bin/env python3
"""
Run inside a generated Python client. Imports every model, then checks that
each Tyk OAS API model (one with an x-tyk-api-gateway field) keeps openapi,
info, paths and x-tyk-api-gateway through from_dict/to_dict. With the empty
OpenAPI3Schema placeholder these models only had x-tyk-api-gateway.
"""
import importlib
import pkgutil
import sys

import openapi_client.models as models

BODY = {
    "openapi": "3.0.3",
    "info": {"title": "Petstore", "version": "1.0.0"},
    "paths": {"/pets": {"get": {"responses": {"200": {"description": "ok"}}}}},
    "x-tyk-api-gateway": {"info": {"name": "Petstore", "state": {"active": True}}},
}
EXPECTED = {"openapi", "info", "paths", "x-tyk-api-gateway"}
EXTENSION_ONLY = {"TykVendorExtension"}  # holds only the vendor extension by design

names = [m.name for m in pkgutil.iter_modules(models.__path__)]
classes = []
for name in names:
    module = importlib.import_module("openapi_client.models." + name)
    for obj in vars(module).values():
        if isinstance(obj, type) and obj.__module__ == module.__name__:
            props = set(getattr(obj, f"_{obj.__name__}__properties", []) or [])
            if "x-tyk-api-gateway" in props and obj.__name__ not in EXTENSION_ONLY:
                classes.append(obj)
print(f"imported {len(names)} models")

if not classes:
    print("❌ no Tyk OAS API model with an x-tyk-api-gateway field found")
    sys.exit(1)

failed = False
for cls in classes:
    kept = set(cls.from_dict(BODY).to_dict())
    ok = EXPECTED <= kept
    failed |= not ok
    print(f"{'✅' if ok else '❌'} {cls.__name__} keeps {sorted(EXPECTED & kept)}" + ("" if ok else f", lost {sorted(EXPECTED - kept)}"))
sys.exit(1 if failed else 0)
