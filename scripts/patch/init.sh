#!/usr/bin/env bash

rm -rf drone
git clone --recursive https://github.com/drone/drone.git drone
cd drone
git checkout $(cat ../VERSION)
go mod vendor
cd --
