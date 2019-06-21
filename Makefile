# General
PKG_LIST_ALL_TESTS	:= $(shell go list ./... | grep -v /vendor | grep -v /test)
GIT_BRANCH			:= $(shell git symbolic-ref HEAD | sed -e 's/^refs\/heads\///')
GIT_LAST_COMMIT		:= $(shell git rev-parse --short HEAD)

# Version
VMAJOR_MINOR 		:= $(or ${VBRANCH}, ${VTAG}, ${GIT_BRANCH})
VBUILD 				:= $(or ${VBUILD}, 0)
VREV 				:= $(or ${VREV}, ${GIT_LAST_COMMIT})
VERSION				:= "$(VMAJOR_MINOR).$(VBUILD).$(shell echo ${VREV} | cut -c 1-8)"

all: help

build: build_alpine build_linux build_osx

build_alpine:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -a -installsuffix cgo -o release/alpine/gcs cmd/main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o release/linux/gcs cmd/main.go

build_osx:
	GOOS=darwin go build -ldflags "-X main.version=$(VERSION)" -o release/osx/gcs cmd/main.go

help:
	@echo 'Usage: '
	@echo ''
	@echo 'make build'
