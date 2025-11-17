#!/bin/bash

set -e

echo "🐹 Generating Go code from OpenAPI..."

# 生成先ディレクトリをクリーンアップ
rm -rf generated/go

# OpenAPI Generatorを使ってGoコードを生成
npx openapi-generator-cli generate \
  -i generated/openapi.yaml \
  -g go \
  -o generated/go \
  --package-name api \
  --git-repo-id mini-notion-api \
  --git-user-id mini-notion \
  --additional-properties=enumClassPrefix=true,generateInterfaces=true

echo "✅ Go code generated at: generated/go/"
