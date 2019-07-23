#!/usr/bin/env sh

PKG_LIST_ALL_TESTS=$(go list ./... | grep -v /vendor | grep -v /test)

# Test!
go test -count=1 -cover $PKG_LIST_ALL_TESTS
