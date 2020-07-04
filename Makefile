VERSION = 0.0.4-dev
BUILD_COMMIT := $(shell git rev-parse HEAD)
BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
DCLI_PKG = github.com/dyweb/gommon/dcli.
DCLI_LDFLAGS = -X $(DCLI_PKG)buildVersion=$(VERSION) -X $(DCLI_PKG)buildCommit=$(BUILD_COMMIT) -X $(DCLI_PKG)buildBranch=$(BUILD_BRANCH) -X $(DCLI_PKG)buildTime=$(BUILD_TIME) -X $(DCLI_PKG)buildUser=$(CURRENT_USER)
FLAGS = $(DCLI_LDFLAGS)
PKGST = ./cmd ./core ./frameworks
PKGS = ./cmd/... ./core... ./frameworks

.PHONY: fmt
fmt:
	goimports -d -l -w $(PKGST)

.PHONY: test
test:
	go test -v -cover $(PKGS)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/bh

.PHONY: clean
clean:
	rm $(shell which bh)

.PHONY: generate
generate:
	gommon generate -v

