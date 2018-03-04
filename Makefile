VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/bhubagent
	go install -ldflags "$(FLAGS)" ./cmd/bhubcentral
	go install -ldflags "$(FLAGS)" ./cmd/bhubctl
# go install -ldflags "$(FLAGS)" ./cmd/bhubdoctor
	go install ./cmd/pingserver
	go install ./cmd/pingclient

.PHONY: clean
clean:
	rm $(shell which bhubagent)
	rm $(shell which bhubcentral)
	rm $(shell which bhubctl)

.PHONY: generate
generate:
	gommon generate -v

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./lib ./pkg

.PHONY: test
test:
	go test -v -cover ./lib/... ./pkg/...