#!/usr/bin/env bash

GENERATOR_IMAGE=openapitools/openapi-generator-cli:v6.0.0

GENERATED="/resources"
OPENAPI_SPEC="/docs/openapi/v1/swagger.yaml"

function generate {
  docker run --rm -v "${PWD}:/local" ${GENERATOR_IMAGE} generate \
      -i "/local${OPENAPI_SPEC}" \
      -g go \
      -o "/local${GENERATED}" \
      --additional-properties packageName="resources" \
      --global-property=models
}

function cleanup_pre {
  rm -rf ".${GENERATED}"
}

cleanup_pre
generate