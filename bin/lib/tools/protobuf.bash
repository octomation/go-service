protogen() {
  rm -rf internal/api
  buf generate
  protoc --twirp_out=internal api/service/v1/service.proto
  go mod tidy
  make format lint test
}
