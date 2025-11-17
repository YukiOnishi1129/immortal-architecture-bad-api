#!/bin/bash

set -euo pipefail

SCRIPT_DIR=$(cd -- "$(dirname "${BASH_SOURCE[0]}")" && pwd)
PROJECT_ROOT=$(cd -- "${SCRIPT_DIR}/.." && pwd)
REPO_ROOT=$(cd -- "${PROJECT_ROOT}/.." && pwd)

GO_TARGET_DIR="${REPO_ROOT}/backend/internal/generated/api"

echo "🐹 Generating Go code from OpenAPI..."

rm -rf "${GO_TARGET_DIR}"

npx openapi-generator-cli generate \
  -i "${PROJECT_ROOT}/generated/openapi.yaml" \
  -g go \
  -o "${GO_TARGET_DIR}" \
  --package-name api \
  --git-repo-id mini-notion-api \
  --git-user-id mini-notion \
  --additional-properties=enumClassPrefix=true,generateInterfaces=true

rm -f "${GO_TARGET_DIR}/go.mod" "${GO_TARGET_DIR}/go.sum"
rm -rf "${GO_TARGET_DIR}/test"

echo "✅ Go code generated at: ${GO_TARGET_DIR}"
