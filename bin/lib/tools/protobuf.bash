protogen() {
  rm -rf api/rpc/v1/*.go api/rpc/v1/*connect/

  buf generate
  protoc \
    --twirp_out=. \
    api/rpc/v1/service.proto

  go mod tidy
  make format lint test
}
