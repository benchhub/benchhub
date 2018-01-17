.PHONY: install
install:
	go install ./cmd/bhubagent
	go install ./cmd/bhubcentral
	go install ./cmd/bhubctl
	go install ./cmd/bhubdoctor
.PHONY: sync-local
sync-local:
	python script/sync_local.py