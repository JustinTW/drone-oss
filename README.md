# Drone OSS

Drone OSS Server with gitlab subgroup support patch.

This repository contains all resources used to build the Docker images available at <https://hub.docker.com/r/justintw/drone-oss>.

### How to build image

```bash
# prepare secret.txt
cp secrets.example.txt secrets.txt

# update your dockerhub creds
vi secrets.txt

# build and push drone-oss image
drone exec \
  --secret-file=secrets.txt \
  --trusted \
  --pipeline build-and-deploy-drone-oss

# build and push drone-nolimit image
drone exec \
  --secret-file=secrets.txt \
  --trusted \
  --pipeline build-and-deploy-drone-nolimit
```

### Changelog

From **v1.7.0** onwards, the tags match the upstream project's tagging. You can now use `garykim/drone-oss:1`, for example.

From **v1.6.3** onwards, a patch is applied before creating the final build. [Patch File](Dockerfile.server.linux.amd64.patch).

### Reference

- [gary-kim/drone-oss](https://github.com/gary-kim/drone-oss)
- [gitlab-workaround-for-subgroups](https://discourse.drone.io/t/gitlab-workaround-for-subgroups/4779)
- [Mrigank11](https://github.com/Mrigank11/go-scm/commit/589b3b440e4a2b9e5d76a336507d11a38123e08e)
