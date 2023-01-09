#!/bin/sh

# usage:
# sh ./run.sh {day} {part}
# sh ./run.sh 01 1

echo go test ./day$1/part2_test.go -input-file=./day$1/test-input

docker run --rm \
    -v "$PWD":/usr/src/myapp \
    -w /usr/src/myapp \
    golang:1.17 \
    go test ./day$1/part2_test.go
