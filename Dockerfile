FROM scratch

LABEL author="OctoLab team <team@octolab.org>"
LABEL org.opencontainers.image.title="Service"
LABEL org.opencontainers.image.description="Template for a typical service written on Go."
LABEL org.opencontainers.image.source="https://github.com/octomation/go-service"
LABEL org.opencontainers.image.licenses="MIT"

COPY server deploy/config.toml /
EXPOSE 3360 8080 8081 8890 8891

ENTRYPOINT ["/server"]
