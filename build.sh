#!/usr/bin/env bash
set -eux

(cd server; go get -t -d -v ./... && go build -v)
