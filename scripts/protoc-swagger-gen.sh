#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
pushd proto

cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)
ibc_go_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/ibc-go/v7)
wasm_dir=$(go list -f '{{ .Dir }}' -m github.com/CosmWasm/wasmd)

proto_dirs=$(find ./ "$cosmos_sdk_dir"/proto "$ibc_go_dir"/proto "$wasm_dir"/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..
# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
npx swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen