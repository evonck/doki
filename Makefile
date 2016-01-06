MAVEN_CACHE_REPO := "$(HOME)/.m2/repository"

.PHONY: all build 

all: check-env docs

check-env:
ifndef DOCKER_HOST
    $(error DOCKER_HOST is undefined)
endif

clean:
	rm -r docs main.js *.json
build:	
	cd ./dokis && make build/container
.PHONY: build

docs: build
	docker run --rm -v `pwd`:/tmp -w /tmp/project boilerplate/dokis-compile ../../dokis -r -f -i / . /tmp/docs
.PHONY: docs
