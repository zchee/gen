.PHONY: all install install-dependencies install-tools lint test test-verbose

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
export ROOT_DIR

ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:) # turn arguments into do-nothing targets
export ARGS

all: install-dependencies install-tools install test

install:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go install ./...
install-dependencies:
	go get -u golang.org/x/tools/imports/...
	go get -u github.com/stretchr/testify/...
	go get -u github.com/termie/go-shutil/...
install-tools:
	# Install linting tools
	go get -u github.com/golang/lint/...
	go get -u github.com/kisielk/errcheck/...

	# Install code coverage tools
	go get -u golang.org/x/tools/cmd/cover/...
	go get -u github.com/onsi/ginkgo/ginkgo/...
	go get -u github.com/modocache/gover/...
	go get -u github.com/mattn/goveralls/...
lint: install
	scripts/lint.sh
test:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s -race ./...
test-verbose:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s -race -v ./...

docker-build:
	docker build --rm -f docker/Dockerfile -t goclang/gen --build-arg LLVM_VERSION=$(LLVM_VERSION) .
docker-test:
	docker run -it --rm -w /go/src/github.com/go-clang/gen -v $(shell pwd):/go/src/github.com/go-clang/gen goclang/gen make ci

ci:
	llvm-config --version
	llvm-config --includedir
	llvm-config --libdir
	make install-dependencies
	make install-tools
	CGO_LDFLAGS="-L`llvm-config --libdir`" go get -u -v -x github.com/go-clang/bootstrap/...
	make lint
	ginkgo -r -cover -skipPackage="testdata"
	gover
	ls -la
