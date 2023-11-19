[![Go Service][social.preview]][preview.config]

[![Coverage][coverage.icon]][coverage.page]
[![Quality][quality.icon]][quality.page]
[![Documentation][docs.icon]][docs.page]
[![CI/CD][build.icon]][build.page]
[![Promo][site.icon]][site.page]
[![Mirror][mirror.icon]][mirror.page]
[![Template][template.icon]][template.page]

# 🧩 Service

Template for a typical service written on Go.

## 🛫 Quick start

<details>
  <summary>Work with Makefile</summary>

```bash
$ make setup
$ make help

$ make find-todos
$ make test lint
$ TIMEOUT=5s make test-with-coverage
```

</details>
<details>
  <summary>Work with Taskfile</summary>

```bash
$ alias run=./Taskfile
$ run refresh
$ run help

$ run docs # === `run docs install -- build -- start`
$ run docs npm ci
$ run docs npm i nextra@latest

$ run tools go generate tools.go
$ run tools golangci-lint --version -- mockgen --version
$ run which goimports golangci-lint govulncheck mockgen
```

</details>
<details>
  <summary>Work with Tools</summary>

```bash
$ make tools
$ source bin/activate

$ which goimports
$ goimports -local $(go list -m) -w ./...
```

</details>
<details>
  <summary>Work with Docker</summary>

```bash
$ make go-1.19 # or go-1.20, etc.
/src# make go-env 2>/dev/null | grep GOVERSION
# GOVERSION:   1.19.10
/src# make test
```

</details>

## 💡 Idea

Define a powerful template that quickly creates a new Go service.
Not only does it provide a starting point for new projects,
but it comes equipped with pre-configured ci/cd and inventory.

```bash
$ server up &
$ client call
```

## 🏆 Motivation

At [OctoLab](https://www.octolab.org/), we want to start new projects faster using best practices
with a predefined structure and focusing on core ideas implementation
rather than wasting time on environment configuration and copying boilerplate code.

## 🤼‍ How to

### Build your own service

1. [Generate][action.generate] a new repository from the template.
2. Clone the repository locally.
3. Update the desired files as needed, e.g., `run init my.new/service`.
4. Write your code and tests.
5. 🚀

### Contribute to the template

1. Read the [contribution guidelines][docs.contrib].
2. [Fork][action.fork] the repository.
3. Make your changes.
4. Send a pull request.
5. 🤗

Before you start, please make sure your changes are in demand.
The best for that is to create [a new discussion][action.discuss],
or if you find an issue, [report it][action.issue] first.

[action.discuss]:   https://github.com/octomation/go-service/discussions/new/choose
[action.fork]:      https://github.com/octomation/go-service/fork
[action.generate]:  https://github.com/octomation/go-service/generate
[action.issue]:     https://github.com/octomation/go-service/issues/new/choose
[docs.contrib]:     https://github.com/octomation/.github/blob/main/.github/CONTRIBUTING.md

## 🎛️ Configuration

### Pre-configured

1. [GitHub Actions](https://github.com/features/actions).
2. [GitHub Pages](https://pages.github.com/).
3. [Dependabot](https://github.com/dependabot).

### Included

1. [Nextra](https://nextra.site/).
2. [Makefiles](https://makefiles.octolab.org/).
3. [Taskfiles](https://taskfiles.octolab.org/).
4. [Go tools](https://github.com/kamilsk/egg):
   - [mockgen](https://github.com/golang/mock)
   - [goimports](https://goimports.octolab.org/)
   - [golangci-lint](https://golangci-lint.octolab.org/)
   - [govulncheck](https://github.com/golang/vuln)
   - [goreleaser](https://goreleaser.com/)
   - [godownloader](https://godownloader.octolab.org/)
   - [protobuf](https://protobuf.dev/)
   - [buf](https://buf.build/)
   - [grpc](https://grpc.io/)
   - [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway/)
   - [twirp](https://twitchtv.github.io/twirp/)

### Optional

1. [Bitbucket](https://bitbucket.org/)[^1].
2. [Codecov](https://about.codecov.io/).
3. [Slack](https://github.com/marketplace/slack-github).
4. [Settings](https://github.com/apps/settings)[^2].
5. [Go Report Card](https://goreportcard.com/).
6. [Shields.io](https://shields.io/).
7. [GitHub Socialify](https://socialify.git.ci/).

[^1]: An alternative for backup could be [GitLab](https://about.gitlab.com/),
[Gogs](https://gogs.io/), or [Gitea](https://gitea.io/).

[^2]: It has been deprecated and will be replaced someday by
[GitHub Actions](https://github.com/octomation/go-module/issues/56).

### Coming soon

1. [Cloudflare Pages](https://pages.cloudflare.com/) (docs hosting).
2. [CodeQL](https://codeql.github.com/) (code scanning).
3. [Graphite](https://graphite.dev/) (git workflow).
4. [Qodana](https://qodana.cloud/) (code quality).
5. [SonarCloud](https://sonarcloud.io/) (code quality).
6. [Vanity URL](https://vanity.octolab.org/) (canonical import path).
7. [Vercel](https://vercel.com/) (docs hosting).

## 🛬 Installation

### Homebrew

```bash
$ brew install octolab/tap/service
```

### Binary

```bash
$ curl -fsSL https://install.octolab.org/octomation/service | sh
# or
$ wget -qO-  https://install.octolab.org/octomation/service | sh
```

### Source

```bash
# use standard go tools
$ go get go.octolab.org/template/service/cmd/client@latest
$ go get go.octolab.org/template/service/cmd/server@latest
# or use egg tool
$ egg tools add go.octolab.org/template/service/cmd/client@latest
$ egg tools add go.octolab.org/template/service/cmd/server@latest
```

### Shell completions

```bash
$ client completion > /path/to/completions/...
# or
$ source <(client completion)
```

## 🏗️ Ecosystem

### Input

- https://github.com/octomation/go-tool
- https://github.com/octopot/forma
- https://github.com/octopot/guard
- https://github.com/octopot/octopus
- https://github.com/octopot/passport
- https://github.com/octopot/tablo

### Impact

- https://github.com/octomation/app
- https://github.com/octomation/install
- https://github.com/octomation/maintainer
- https://github.com/octomation/makefiles
- https://github.com/octomation/taskfiles
- https://github.com/octomation/vanity

### Output

- https://github.com/withsparkle/service

<p align="right">made with ❤️ for everyone by <a href="https://www.octolab.org/">OctoLab</a></p>

[social.preview]:   https://cdn.octolab.org/repo/go-service.png
[preview.config]:   https://socialify.git.ci/octomation/go-service?description=1&font=Raleway&language=1&name=1&owner=1&pattern=Circuit%20Board&theme=Light
[preview.fallback]: https://socialify.git.ci/octomation/go-service/image?description=1&font=Raleway&language=1&name=1&owner=1&pattern=Circuit%20Board&theme=Light

[awesome.icon]:     https://awesome.re/mentioned-badge.svg
[awesome.page]:     https://awesome-go.com/project-layout/
[coverage.icon]:    https://codecov.io/gh/octomation/go-service/branch/main/graph/badge.svg
[coverage.page]:    https://codecov.io/gh/octomation/go-service
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org/template/service
[quality.page]:     https://goreportcard.com/report/go.octolab.org/template/service
[docs.icon]:        https://pkg.go.dev/badge/go.octolab.org/template/service.svg
[docs.page]:        https://pkg.go.dev/go.octolab.org/template/service
[build.icon]:       https://img.shields.io/badge/ci%2Fcd-GitHub%20Actions-brightgreen
[build.page]:       https://github.com/octomation/go-service/actions
[site.icon]:        https://img.shields.io/badge/site-GitHub%20Pages-brightgreen
[site.page]:        https://go-service.octolab.org/
[mirror.icon]:      https://img.shields.io/badge/mirror-Bitbucket-blue
[mirror.page]:      https://bitbucket.org/kamilsk/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue
[template.page]:    https://github.com/octomation/go-service
