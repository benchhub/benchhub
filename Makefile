VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

.PHONY: install
install:
	go install -ldflags "$(FLAGS)" ./cmd/bhubagent
	go install -ldflags "$(FLAGS)" ./cmd/bhubcentral
	go install -ldflags "$(FLAGS)" ./cmd/bhubagentctl
	go install -ldflags "$(FLAGS)" ./cmd/bhubctl
# go install -ldflags "$(FLAGS)" ./cmd/bhubdoctor
	go install ./cmd/pingserver
	go install ./cmd/pingclient
	go install ./lib/waitforit/cmd/waitforit

.PHONY: clean
clean:
	rm $(shell which bhubagent)
	rm $(shell which bhubcentral)
	rm $(shell which bhubctl)

.PHONY: loc
loc:
	cloc --exclude-dir=vendor,.idea,playground,vagrant,node_modules,bhpb, --exclude-list-file=script/cloc_exclude.txt .

.PHONY: generate
generate:
	gommon generate -v

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./lib ./pkg

.PHONY: test
test:
	go test -v -cover ./lib/... ./pkg/...

.PHONY: package
package: install
	cp $(shell which bhubagent) .
	cp $(shell which bhubcentral) .
	cp $(shell which bhubctl) .
	cp $(shell which bhubagentctl) .
	cp $(shell which pingserver) .
	cp $(shell which pingclient) .
	cp $(shell which waitforit) .
	zip bhubagent-$(VERSION).zip bhubagent
	zip bhubcentral-$(VERSION).zip bhubcentral
	zip bhubctl-$(VERSION).zip bhubctl
	zip bhubagentctl-$(VERSION).zip bhubagentctl
	zip pingserver-$(VERSION).zip pingserver
	zip pingclient-$(VERSION).zip pingclient
	zip waitforit-$(VERSION).zip waitforit
	rm bhubagent
	rm bhubcentral
	rm bhubctl
	rm bhubagentctl
	rm pingserver
	rm pingclient
	rm waitforit
