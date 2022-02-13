#!/bin/bash

oss=(darwin linux)
archs=(amd64)

for os in ${oss[@]}
    do
    for arch in ${archs[@]}
    do
        CGO_ENABLED=0 GOOS=${os} GOARCH=${arch} go build -ldflags="-s -w" -o bin/action-${os}-${arch} main.go
    done
done
