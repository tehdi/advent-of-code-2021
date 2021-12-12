#!/bin/sh

# usage:
# sh ./run.sh {day} {part}
# sh ./run.sh 01 1

echo go run ./day$1/part$2.go -input-file=./day$1/test-input

docker run --rm \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    golang:1.17 \
    go run ./day$1/part$2.go -input-file=./day$1/test-input
