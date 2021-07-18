#!/usr/bin/env bash

rm -rf drone

drone exec --secret-file=secrets.txt --trusted --pipeline build-and-deploy-drone-nolimit
