# should be run on host for sync binary etc.
.PHONY: copy
copy: copy-bin copy-config

.PHONY: copy-bin
copy-bin:
	cp $(shell which bhubagent) .
	cp $(shell which bhubcentral) .
	cp $(shell which bhubctl) .
	cp $(shell which pingserver) .
	cp $(shell which pingclient) .

.PHONY: copy-config
copy-config:
	cp ../../bhubagent.yml .
	cp ../../bhubcentral.yml .