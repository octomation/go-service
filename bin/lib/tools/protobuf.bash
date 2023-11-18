protogen() {
  rm -rf api/rpc/v1/*.go api/rpc/v1/*connect/
  local opts openapi=(
    include_description=true
    single_file=true
    yaml=true
  )
  IFS=','
  opts=$(printf "%s" "${openapi[*]}")

  buf generate
  protoc \
    --openapi_out="${opts}:api/openapi" \
    --twirp_out=. \
    api/rpc/v1/service.proto

  go mod tidy
  make format lint test

  # experimental
  local spec=api/openapi/spec.v3.yaml
  yq ea '. as $item ireduce ({}; . * $item)' \
    api/openapi/{meta,openapiv3}.yaml \
    >"${spec}"
  if gum confirm 'Use openapi hack?'; then
    yq ea '. as $item ireduce ({}; . * $item)' \
      api/openapi/{openapiv3,meta}.yaml \
      >"${spec}"
    yq -i 'sort_keys(.)' "${spec}"
  fi
  rm -f api/openapi/openapiv3.yaml
}
