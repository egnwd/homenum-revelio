#!/usr/bin/env bash
set -eux

(cd client; webpack)
(cd server; go get -t -d -v ./... && go build -v)
