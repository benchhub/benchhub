VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)
PKGST = ./cmd ./pkg
PKGS = ./cmd/... ./pkg...

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

