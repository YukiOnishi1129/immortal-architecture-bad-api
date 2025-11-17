GOBIN ?= $(shell go env GOPATH)/bin
GOLANGCI_LINT ?= $(GOBIN)/golangci-lint
GOCACHE ?= $(CURDIR)/.gocache
GOMODCACHE ?= $(CURDIR)/.gomodcache

.PHONY: lint tidy

lint:
	@GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) $(GOLANGCI_LINT) run ./...

tidy:
	@GOCACHE=$(GOCACHE) GOMODCACHE=$(GOMODCACHE) go mod tidy
