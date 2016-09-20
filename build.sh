#!/usr/bin/env bash
set -eux

(cd client; npm install && gulp && rm -rf ../bin/dist && mv -f dist/ ../bin/)
(cd server; go get -t -d -v ./... && go build -v && mv server ../bin/server)
