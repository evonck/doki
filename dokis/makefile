# Old-skool build tools.
#
# Targets (see each target for more information):
#   build/container:	builds the Docker image used to compile all golang code
#   build:        		builds binaries, placing each next to it's respective main pkg
#   test: 		    		runs tests
#   lint:    				  lints the source tree
#   install: 					builds, tests, then copies the resulting binary to $GOPATH/bin/
#   dockerize:    		builds, tests, then makes a Docker image for each binary
#	  clean:			  		removes build artifacts (aka binaries)
#	  clean-all:	  		cleans, then removes the artifact for build/con

SHELL := /bin/bash
ProjectName := dokis

# some terminal color escape codes
LIGHT_GREEN := $(shell echo -e "\033[1;32m")
NC := $(shell echo -e "\033[0m") # No Color

GO ?= go

# the "root" pkg contained in this project.  If the project contains multiple
# binaries, each binary's main pkg will be in a subdir of SRC_ROOT
SRC_ROOT= github.com/evonck/doki/

.DEFAULT_GOAL := build

# Builds the docker image that we'll use to compile all subsequent golang code
# touch: http://www.gnu.org/software/make/manual/make.html#Empty-Targets
build/container: buildall
	@echo "${LIGHT_GREEN}building Docker image: boilerplate/$(ProjectName)-compile... ${NC}"
	@docker build -t boilerplate/$(ProjectName)-compile . > /dev/null

clean:
	rm -f  $(ProjectName)
	rm -rf ./bin

clean-all: clean
	@docker rm -v $(ProjectName)
	@docker rmi boilerplate/$(ProjectName)-compile


# The release number & build date are stamped into the binary.
build:	
	go build -v -i -o $(ProjectName)
.PHONY: build


# Compile the binary for all available OSes and ARCHes.
# go get https://github.com/mitchellh/gox
buildall: OSARCH += -osarch "darwin/amd64 linux/amd64 linux/386"
buildall:
	@echo "${LIGHT_GREEN}building Docker in /bin ${NC}"
	gox $(OSARCH) -output "bin/$(ProjectName)_{{.OS}}_{{.Arch}}"
.PHONY: buildall


# runs any tests inside a Docker container
test: 
	@echo "${LIGHT_GREEN}running tests for $(ProjectName) ${NC}"
	godep go test -v ./...
.PHONY: test

# runs any check inside a Docker container
check: build/container
	@echo "${LIGHT_GREEN}running check for $(ProjectName)...${NC}"
	@docker run --rm \
		--name $(ProjectName) \
		-v "$$PWD":"/srv/src/${SRC_ROOT}" \
		-w "/srv/src/${SRC_ROOT}" \
		boilerplate/$(ProjectName)-compile \
		golint ./...
.PHONY: check


install: build test
	@echo "${LIGHT_GREEN}copying binary to ${GOPATH}/bin/...${NC}"
	cp "$$PWD"/$(ProjectName) $${GOPATH}/bin
.PHONY: install


# Build a linux-compatible binary and a docker image that uses the binary as it's entrypoint
dockerize: build/container
	@echo "${LIGHT_GREEN}building /$(ProjectName) binary for inclusion in Docker image...${NC}"
	@docker run --rm -t boilerplate/$(ProjectName)-compile
.PHONY: dockerize

dockerCompose:  build/container
	@echo "${LIGHT_GREEN}launch /$(ProjectName) docker image and all of it's dependency ${NC}"
	@docker-compose up


