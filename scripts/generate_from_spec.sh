#!/bin/bash

set -xe pipefail

openapi_spec_path="./api/openapi-spec/apimanager.yaml"

for tag in Actions Downloads Gateways Organizations Plugins "Policy Definitions" Roles Search System Users; do

  lowercase_tag="${tag,,}"

  # replace spaces with underscore
  lowercase_tag=${lowercase_tag// /_}

  echo normal: "${tag}", lowercase: "${lowercase_tag}"

  dir_path="./internal/app/apimanager/${lowercase_tag}"

  # Create dir path if it doesn't exist
  mkdir -p "${dir_path}"

  # Clean up directory in case it already exists to remove any unwanted files
  rm -rf "${dir_path}/*"

  # Generate Types
  oapi-codegen -generate types -include-tags="${tag}" -package="${lowercase_tag}" ${openapi_spec_path} > "${dir_path}/types.gen.go"

  # Generate Server Interface
  oapi-codegen -generate server,spec -include-tags="${tag}" -package="${lowercase_tag}" ${openapi_spec_path} > "${dir_path}/server.gen.go"
done

