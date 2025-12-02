---
id: 66
database_id: 1402337161
node_id: I_kwDOD0sMFc5TlfeJ
status: closed
title: "ci/cd: dependabot: change commit message"
labels: ["type: improvement","scope: inventory","impact: medium","effort: easy"]
url: https://github.com/octomation/go-service/issues/66
created_at: 2022-10-09T17:10:44Z
updated_at: 2023-11-18T09:23:56Z
---

# ci/cd: dependabot: change commit message

**Motivation:** to be precise.

- `directory: /` -> `deps:`
- `directory: /docs/` -> `deps: docs:` or `docs: deps:`
- `directory: /tools/` -> `deps: tools:` or `tools: deps:`
- `github-actions` -> `deps: ci/cd:` or `ci/cd: deps:`

See the docs https://docs.github.com/en/code-security/dependabot/dependabot-version-updates/configuration-options-for-the-dependabot.yml-file#commit-message.
Don't forget about https://www.conventionalcommits.org/en/v1.0.0/.

Related to octomation/go-module/issues/65.
