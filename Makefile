.PHONY: install
install:
	go install ./cmd/bhubagent
	go install ./cmd/bhubcentral
	go install ./cmd/bhubctl
	go install ./cmd/bhubdoctor

.PHONY: fmt
fmt:
	gofmt -d -l -w ./cmd ./lib ./pkg

.PHONY: test
test:
	go test -v -cover ./lib/... ./pkg/...