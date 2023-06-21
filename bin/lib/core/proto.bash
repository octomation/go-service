#!/usr/bin/env bash
# shellcheck source=../utils/print.bash # @error @fatal

#eval "$(run gateway.proto)"
gateway.proto() {
  cat <<EOF
    rm -rf api/vendor/protoc-gen-openapiv2/options/*
    pushd /tmp || return 1
    clone git@github.com:grpc-ecosystem/grpc-gateway.git#v2.7.3
    popd || return 1
    popd || return 1
    find /tmp/grpc-gateway/protoc-gen-openapiv2/options \
      -maxdepth 1 \
      -name '*.proto' \
      -exec cp -f {} api/vendor/protoc-gen-openapiv2/options \;
    rm -rf /tmp/grpc-gateway
EOF
}

#eval "$(run google.proto)"
google.proto() {
  cat <<EOF
    rm -rf api/vendor/google/protobuf/*
    pushd /tmp || return 1
    clone git@github.com:protocolbuffers/protobuf.git#v3.19.4
    popd || return 1
    popd || return 1
    find /tmp/protobuf/src/google/protobuf \
      -maxdepth 1 \
      -name '*.proto' \
      -exec cp -f {} api/vendor/google/protobuf \;
    rm -rf /tmp/protobuf
EOF
}
