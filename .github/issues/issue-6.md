---
id: 6
database_id: 848607736
node_id: MDU6SXNzdWU4NDg2MDc3MzY=
status: open
title: "dist: docker: compare distroless image with alpine"
labels: ["type: improvement","scope: inventory","impact: medium","effort: medium"]
url: https://github.com/octomation/go-service/issues/6
created_at: 2021-04-01T16:18:45Z
updated_at: 2023-11-18T09:32:08Z
---

# dist: docker: compare distroless image with alpine

**Motivation:** "Distroless" images contain only your application and its runtime dependencies. They do not contain package managers, shells, or any other programs you would expect to find in a standard Linux distribution.

See https://github.com/GoogleContainerTools/distroless.

**Inspiration:** https://github.com/itchyny/gojq/commit/80d1ed72b977163175f7655117bececdfecf4b34.
