

PLATFORMS := darwin windows linux 
OS ?= darwin
VERSION ?= latest
BUMP ?= minor
BINARY := dockument
os = $(word 1, $@)

release: 
	mkdir -p release 

.PHONY: $(PLATFORMS) 
$(PLATFORMS): 
	@- cd cmd && GOOS=$(os) GOARCH=amd64 go build -o ../release/$(BINARY)-$(VERSION)-$@-amd64

.PHONY: install 
install: releases
ifeq ($(OS),darwin)
	cp release/$(BINARY)-$(VERSION)-darwin-amd64 $(GOPATH)/bin/$(BINARY)
else ifeq ($(OS),linux)
	cp release/$(BINARY)-$(VERSION)-linux-amd64 $(GOPATH)/bin/$(BINARY)
else 
	cp release/$(BINARY)-$(VERSION)-windows-amd64 $(GOPATH)/bin/$(BINARY)
endif



.PHONY: releases 
releases: release darwin windows linux

.PHONY: dry
dry: 
	@- cd cmd && CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 go build -ldflags="-X main.documentation=" -o ../$(BINARY)

.PHONY: doc 
doc: dry 
	@-./$(BINARY)  > /dev/null  2>&1 || true  

.PHONY: bumpversion 
bumpversion: 
	@- chmod +x versionutils/bumpversion.sh
	@- ./versionutils/bumpversion.sh $(PWD)/VERSION $(BUMP)