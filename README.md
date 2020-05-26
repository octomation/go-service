> # 🧩 Service
>
> Template for typical Go service.

[![Build][build.icon]][build.page]
[![Template][template.icon]][template.page]

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
$ curl -sSfL https://:install.sh | sh
# or
$ wget -qO-  https://:install.sh | sh
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

### Bash and Zsh completions

```bash
$ :binary completion bash > /path/to/bash_completion.d/:binary.sh
$ :binary completion zsh  > /path/to/zsh-completions/_:binary.zsh
# or autodetect
$ source <(:binary completion)
```

> See `kubectl` [documentation](https://kubernetes.io/docs/tasks/tools/install-kubectl/#enabling-shell-autocompletion).

## 🤲 Outcomes

...

---

made with ❤️ for everyone

[build.page]:       https://travis-ci.com/:owner/:repository
[build.icon]:       https://travis-ci.com/:owner/:repository.svg?branch=master
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0?r=0b753cbf767346f5a6fd51194829a2f3
[promo.page]:       https://github.com/:owner/:repository
[template.page]:    https://github.com/octomation/go-service
[template.icon]:    https://img.shields.io/badge/template-go--service-blue

[egg]:              https://github.com/kamilsk/egg
