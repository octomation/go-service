> # 🧩 Service
>
> Template for typical Go service.

[![Build][build.icon]][build.page]
[![Documentation][docs.icon]][docs.page]
[![Quality][quality.icon]][quality.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]
[![Mirror][mirror.icon]][mirror.page]

## 💡 Idea

...

A full description of the idea is available [here][design.page].

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install :owner/tap/:binary
```

### Binary

```bash
$ curl -fsSL https://raw.githubusercontent.com/octomation/go-service/main/bin/install | sh
# or
$ wget -qO-  https://raw.githubusercontent.com/octomation/go-service/main/bin/install | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get github.com/:owner/:repository/cmd/client@:version
$ go get github.com/:owner/:repository/cmd/server@:version
# or use egg tool
$ egg tools add github.com/:owner/:repository/cmd/client@:version
$ egg tools add github.com/:owner/:repository/cmd/server@:version
```

> [egg][] is an `extended go get`.

### Shell completions

```bash
$ :binary completion > /path/to/completions/...
# or
$ source <(:binary completion)
```

## 🤲 Outcomes

...

<p align="right">made with ❤️ for everyone</p>

[awesome.icon]:     https://awesome.re/mentioned-badge.svg
[build.page]:       https://github.com/:owner/:repository/actions/workflows/ci.yml
[build.icon]:       https://github.com/octomation/go-service/actions/workflows/ci.yml/badge.svg
[coverage.page]:    https://codeclimate.com/github/:owner/:repository/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/4c01eaeb09061b2ced6b/test_coverage
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0?r=0b753cbf767346f5a6fd51194829a2f3
[docs.page]:        https://pkg.go.dev/:module/:version
[docs.icon]:        https://img.shields.io/badge/docs-pkg.go.dev-blue
[mirror.page]:      https://bitbucket.org/kamilsk/go-service
[mirror.icon]:      https://img.shields.io/badge/mirror-bitbucket-blue
[promo.page]:       https://github.com/:owner/:repository
[quality.page]:     https://goreportcard.com/report/:module
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg
