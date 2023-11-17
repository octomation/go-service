# shellcheck source=../git/core.bash # @root

@up() {
  @compose up -d
}

@down() {
  @compose down
}

@destroy() {
  @compose down -v
}

@compose() {
  docker compose \
    --env-file .env \
    --project-directory "$(@root)" \
    -f env/docker-compose.yml \
    -p sparkle \
    "${@}"
}
