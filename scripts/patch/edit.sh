#!/usr/bin/env bash

# gitlab subgroup support
# ref: https://discourse.drone.io/t/gitlab-workaround-for-subgroups/4779

cp drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/util.go \
  drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/util-patched.go
patch drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/util-patched.go \
  < patch/go-scm/gitlab-util.go

cp drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/repo.go \
  drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/repo-patched.go
patch drone/vendor/github.com/drone/go-scm/scm/driver/gitlab/repo-patched.go \
  < patch/go-scm/gitlab-repo.go

cp drone/vendor/github.com/drone/go-scm/scm/util.go \
  drone/vendor/github.com/drone/go-scm/scm/util-patched.go
patch drone/vendor/github.com/drone/go-scm/scm/util-patched.go \
  < patch/go-scm/util.go