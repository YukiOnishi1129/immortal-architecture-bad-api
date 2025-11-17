#!/bin/bash

set -e

echo "📄 Generating OpenAPI YAML from TypeSpec..."

# TypeSpecコンパイル
tsp compile typespec --emit @typespec/openapi3

echo "✅ OpenAPI YAML generated at: generated/openapi.yaml"
