#!/bin/bash

set -e

echo "📘 Generating TypeScript code from OpenAPI..."

# 生成先ディレクトリをクリーンアップ
rm -rf generated/typescript

# OpenAPI Generatorを使ってTypeScriptコードを生成
npx openapi-generator-cli generate \
  -i generated/openapi.yaml \
  -g typescript-axios \
  -o generated/typescript \
  --additional-properties=withSeparateModelsAndApi=true,apiPackage=client,modelPackage=models,supportsES6=true,useSingleRequestParameter=true

echo "✅ TypeScript code generated at: generated/typescript/"
