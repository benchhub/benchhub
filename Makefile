VERSION = 0.0.4-dev
# Build Info
BUILD_COMMIT := $(shell git rev-parse HEAD)
BUILD_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
# Go
GO = go
DCLI_PKG = github.com/dyweb/gommon/dcli.
DCLI_LDFLAGS = -X $(DCLI_PKG)buildVersion=$(VERSION) -X $(DCLI_PKG)buildCommit=$(BUILD_COMMIT) -X $(DCLI_PKG)buildBranch=$(BUILD_BRANCH) -X $(DCLI_PKG)buildTime=$(BUILD_TIME) -X $(DCLI_PKG)buildUser=$(CURRENT_USER)
FLAGS = $(DCLI_LDFLAGS)
PKGST = ./cmd ./core ./frameworks ./lib ./runtimes
PKGS = $(addsuffix ...,$(PKGST))

.PHONY: fmt
fmt:
	gommon format -d -l -w $(PKGST)

.PHONY: test
test:
	$(GO) test -v -cover $(PKGS)

.PHONY: install
install: install-generator
	$(GO) install -ldflags "$(FLAGS)" ./cmd/bh

install-generator:
	$(GO) install -ldflags "$(FLAGS)" ./cmd/bhgen

uninstall:
	rm $(shell which bh)
	rm $(shell which bhgen)

.PHONY: clean
clean:
	bhgen schema clean
	rm -rf bhpb/*.pb.go

.PHONY: generate
generate: gen-proto gen-schema

gen-proto: install-generator
	gommon generate -v

gen-schema: install-generator
	bhgen schema generate
	$(GO) run build/generated/tqbuilder/ddl/main.go
