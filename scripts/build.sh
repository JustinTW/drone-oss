#!/usr/bin/env bash
if [[ "$1" == "nolimit" ]]; then
  go build -mod vendor --tags "nolimit" -ldflags "-extldflags \"-static\"" -o release/linux/amd64/drone-server github.com/drone/drone/cmd/drone-server
  exit 0
fi

go build -mod vendor --tags "oss nolimit" -ldflags "-extldflags \"-static\"" -o release/linux/amd64/drone-server github.com/drone/drone/cmd/drone-server
