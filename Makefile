# General
GIT_BRANCH			:= $(shell git symbolic-ref HEAD | sed -e 's/^refs\/heads\///')
GIT_LAST_COMMIT		:= $(shell git rev-parse --short HEAD)

# Version
VMAJOR_MINOR 		:= $(or ${VBRANCH}, ${VTAG}, ${GIT_BRANCH})
VBUILD 				:= $(or ${VBUILD}, 0)
VREV 				:= $(or ${VREV}, ${GIT_LAST_COMMIT})
VERSION				:= "$(VMAJOR_MINOR).$(VBUILD).$(shell echo ${VREV} | cut -c 1-8)"

default: all

all: clean test bin dist

clean:
	@rm -rf target

test:
	@sh -c "'$(PWD)/scripts/unit.test.sh'"

bin:
	@VERSION=$(VERSION) sh -c "'$(PWD)/scripts/build.sh'"

dist:
	@VERSION=$(VERSION) sh -c "'$(PWD)/scripts/dist.sh'"

publish:
	@VERSION=$(VERSION) sh -c "'$(PWD)/scripts/publish.sh'"

.NOTPARALLEL:

.PHONY: clean test bin dist publish
