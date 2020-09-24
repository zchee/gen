.DEFAULT_GOAL = dev

.PHONY: all install install-dependencies install-tools lint test test-full test-verbose

export ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

export CC := clang
export CXX := clang++

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

	CGO_LDFLAGS="-L`llvm-config --libdir`" go get github.com/go-clang/bootstrap/...
install-tools:
	# Install linting tools
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck/...

	# Install code coverage tools
	go get -u golang.org/x/tools/cmd/cover/...
	go get -u github.com/onsi/ginkgo/ginkgo/...
	go get -u github.com/modocache/gover/...
	go get -u github.com/mattn/goveralls/...
lint: install
	$(ROOT_DIR)/scripts/lint.sh
test:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s ./...
test-full:
	$(ROOT_DIR)/scripts/test-full.sh
test-verbose:
	CGO_LDFLAGS="-L`llvm-config --libdir`" go test -timeout 60s -v ./...

dev/%:
	# CGO_CFLAGS="-fPIC -Wall -W -Wno-unused-variable -Wno-language-extension-token -Wno-unused-parameter -Wwrite-strings -Wmissing-field-initializers -pedantic -Wno-long-long -Wcovered-switch-default -Wdelete-non-virtual-dtor -Werror=date-time -O3 -DNDEBUG -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS -Wno-deprecated-declarations" CGO_LDFLAGS="-L$(shell /opt/llvm/$*/bin/llvm-config --libdir) -Wl,-rpath,$(shell /opt/llvm/$*/bin/llvm-config --libdir)" go-install-pkg || true
	CGO_CFLAGS="-fPIC -Wall -W -Wno-unused-variable -Wno-language-extension-token -Wno-unused-parameter -Wwrite-strings -Wmissing-field-initializers -pedantic -Wno-long-long -Wcovered-switch-default -Wdelete-non-virtual-dtor -Werror=date-time -O3 -DNDEBUG -D__STDC_CONSTANT_MACROS -D__STDC_FORMAT_MACROS -D__STDC_LIMIT_MACROS -Wno-deprecated-declarations" CGO_LDFLAGS="-L$(shell /opt/llvm/$*/bin/llvm-config --libdir) -Wl,-rpath,$(shell /opt/llvm/$*/bin/llvm-config --libdir)" go build -v -x -o go-clang-gen-$* ./cmd/go-clang-gen
	@cd ../clang-v3.9-test/clang; \
		rm -rf clang-c '*_gen.go'; \
		PATH=/opt/llvm/$*/bin:$$PATH ../../gen/go-clang-gen-$* /opt/llvm/$*/lib/clang
