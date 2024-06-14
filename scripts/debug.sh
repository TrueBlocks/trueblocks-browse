#!/usr/bin/env bash

export CGO_LDFLAGS="-framework CoreServices -framework UniformTypeIdentifiers"

mkdir -p build/bin
go build -gcflags "all=-N -l" -tags "dev,devtools" -o build/bin/out
ls -la build/bin/out
