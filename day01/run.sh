#!/bin/sh

docker run --rm \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    golang:1.17 \
    go run $1
