@build() {
  local ldflags
  ldflags="-s -w -X main.commit=$(git rev-parse HEAD)"
  ldflags="$ldflags -X main.date=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
  ldflags="$ldflags -X main.version=$(git describe --tags 2>/dev/null || echo dev)"

  docker build \
    --build-arg ldflags="${ldflags}" \
    --tag ghcr.io/octomation/go-service:latest \
    .
}

@buildx() {
  if ! docker buildx use builder 2>/dev/null; then
    docker buildx create --name builder --use
    docker buildx inspect --bootstrap
  fi
  docker buildx build \
    --platform linux/amd64,linux/arm64 \
    --tag ghcr.io/octomation/go-service:latest \
    .
}
