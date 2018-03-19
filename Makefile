BIN_DIR := $(GOPATH)/bin
PLATFORMS := darwin windows linux 
OS ?= darwin
VERSION ?= latest
BUMP ?= minor
BINARY := dockument
GOVERAGE := $(BIN_DIR)/goverage 
GOMETALINTER := $(BIN_DIR)/gometalinter
PKGS := $(shell go list ./... | grep -v /vendor)
os = $(word 1, $@)


.PHONY: test 
test: $(GOVERAGE) 
	go test $(PKGS)
	# goverage -race -coverprofile=coverage.out ./...
    # go tool cover -html=coverage.out
	
$(GOMETALINTER):
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install &> /dev/null


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

.PHONY: lint 
lint: $(GOMETALINTER)
	gometalinter ./... --vendor --errors

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