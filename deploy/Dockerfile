FROM golang:1.21-alpine as builder

ARG ldflags

ENV CGO_ENABLED=0
ENV LDFLAGS=${ldflags:-'-s -w'}

WORKDIR /service

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build \
  -o server \
  -trimpath \
  -ldflags "${LDFLAGS}" \
  ./cmd/server


FROM scratch

LABEL author="OctoLab team <team@octolab.org>"
LABEL org.opencontainers.image.title="Service"
LABEL org.opencontainers.image.description="Template for a typical service written on Go."
LABEL org.opencontainers.image.source="https://github.com/octomation/go-service"
LABEL org.opencontainers.image.licenses="MIT"

COPY --from=builder /service/server /service/deploy/config.toml /
EXPOSE 3360 8080 8081 8890 8891

ENTRYPOINT ["/server"]
CMD ["run"]
