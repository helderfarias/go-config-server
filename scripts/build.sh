#!/usr/bin/env bash

# Global
XC_OS=${XC_OS:-"$(go env GOOS)"}
XC_ARCH=${XC_ARCH:-"$(go env GOARCH)"}
LD_FLAGS=${LD_FLAGS:-"-X main.version=$VERSION"}
CGO_ENABLED=${CGO_ENABLED:-0}

# Deps!
echo "==> Download deps...: $XC_OS/$XC_ARCH "
go mod download

# Build!
echo
echo "==> Building........: $XC_OS/$XC_ARCH"
CGO_ENABLED=$CGO_ENABLED GOOS=$XC_OS GOARCH=$XC_ARCH go build -ldflags "$LD_FLAGS" -a -installsuffix cgo -o target/gcs_$XC_OS cmd/main.go

# Done!
echo
echo "==> Results.........: "
ls -hl target
