---
code: null
id: MDU6SXNzdWU4NDg2MDc3MzY=
databaseId: 848607736
number: 6
url: https://github.com/octomation/go-service/issues/6
title: "dist: docker: compare distroless image with alpine"
issueType: null
labels:
  - "type: improvement"
  - "scope: inventory"
  - "impact: medium"
  - "effort: medium"
assignees:
  - kamilsk
milestone: "[[milestone-1]]"
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2021-04-01T16:18:45Z
updatedAt: 2023-11-18T09:32:08Z
lastEditedAt: 2022-05-25T08:39:39Z
closedAt: null
---

# dist: docker: compare distroless image with alpine

**Motivation:** "Distroless" images contain only your application and its runtime dependencies. They do not contain package managers, shells, or any other programs you would expect to find in a standard Linux distribution.

See https://github.com/GoogleContainerTools/distroless.

**Inspiration:** https://github.com/itchyny/gojq/commit/80d1ed72b977163175f7655117bececdfecf4b34.

<!-- 2022-06-26T13:00Z https://github.com/octomation/go-service/issues/6#issuecomment-1166525676
https://medium.com/inside-sumup/stop-using-alpine-docker-images-fbf122c63010
-->
