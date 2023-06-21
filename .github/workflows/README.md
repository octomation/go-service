[![Go Service][social.preview]][preview.config]

# 🧩 Service

GitHub Actions workflows.

✅ Slack notification is available for all workflows.
Just add `SLACK_WEBHOOK` secret to your repository.

## [Cache invalidation](caches.yml)

[![Status][caches.icon]][caches.page]

Invalidates caches of GitHub Actions workflows.
Read more about caches https://docs.github.com/en/actions/using-workflows/caching-dependencies-to-speed-up-workflows.


## [Continuous delivery](cd.yml)

[![Status][cd.icon]][cd.page]

Delivery binaries using [GoReleaser](https://goreleaser.com).


## [Continuous integration](ci.yml)

[![Status][ci.icon]][ci.page]

Runs linter and tests on different Go environments.


## [Dependabot at weekends](deps.yml)

[![Status][deps.icon]][deps.page]

🚧 Experimental feature, still under development. Read more

- https://github.com/orgs/community/discussions/15901
- https://github.com/dependabot/dependabot-core/issues/2980#issuecomment-760586879
- https://github.com/octomation/go-module/issues/79
- https://github.com/dependabot/cli
- https://github.com/github/dependabot-action

⚠️ Required:

- `DEPENDABOT_TOKEN`, GitHub personal access tokens with scope: gist.


## [Documentation delivery](docs.yml)

[![Status][docs.icon]][docs.page]

Builds and deploys documentation to GitHub Pages.
[💬 Discussion](https://github.com/under-the-hood/docs/discussions/2).


## [Workflow invalidation](runs.yml)

[![Status][runs.icon]][runs.page]

⚠️ Removes old workflow runs. Please, use it carefully.


## [Issue invalidation](stale.yml)

[![Status][stale.icon]][stale.page]

Checks for stale issues and pull requests.


## [Tools validation](tools.yml)

[![Status][tools.icon]][tools.page]

Checks tools for consistency.

<p align="right">made with ❤️ for everyone by OctoLab</p>

[social.preview]:   https://socialify.git.ci/octomation/go-service/image?description=1&font=Raleway&language=1&name=1&owner=1&pattern=Circuit%20Board&theme=Light
[preview.config]:   https://socialify.git.ci/octomation/go-service?description=1&font=Raleway&language=1&name=1&owner=1&pattern=Circuit%20Board&theme=Light
[preview.fallback]: https://repository-images.githubusercontent.com/256576533/bf982df3-0bbc-44a1-8f55-e819eeb39a1c

[caches.icon]:      https://github.com/octomation/go-service/actions/workflows/caches.yml/badge.svg
[caches.page]:      https://github.com/octomation/go-service/actions/workflows/caches.yml
[cd.icon]:          https://github.com/octomation/go-service/actions/workflows/cd.yml/badge.svg
[cd.page]:          https://github.com/octomation/go-service/actions/workflows/cd.yml
[ci.icon]:          https://github.com/octomation/go-service/actions/workflows/ci.yml/badge.svg
[ci.page]:          https://github.com/octomation/go-service/actions/workflows/ci.yml
[deps.icon]:        https://github.com/octomation/go-service/actions/workflows/deps.yml/badge.svg
[deps.page]:        https://github.com/octomation/go-service/actions/workflows/deps.yml
[docs.icon]:        https://github.com/octomation/go-service/actions/workflows/docs.yml/badge.svg
[docs.page]:        https://github.com/octomation/go-service/actions/workflows/docs.yml
[runs.icon]:        https://github.com/octomation/go-service/actions/workflows/runs.yml/badge.svg
[runs.page]:        https://github.com/octomation/go-service/actions/workflows/runs.yml
[stale.icon]:       https://github.com/octomation/go-service/actions/workflows/stale.yml/badge.svg
[stale.page]:       https://github.com/octomation/go-service/actions/workflows/stale.yml
[tools.icon]:       https://github.com/octomation/go-service/actions/workflows/tools.yml/badge.svg
[tools.page]:       https://github.com/octomation/go-service/actions/workflows/tools.yml
