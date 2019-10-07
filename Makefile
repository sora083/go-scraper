export GO111MODULE=on

NAME := go-scraper
VERSION := 1.0.0
REVISION := $(shell git rev-parse --short HEAD)

#BUILD := $(shell git rev-parse --short HEAD)
#PROJECTNAME := $(shell basename "$(PWD)")

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""
NOVENDOR := $(shell go list ./... | grep -v vendor)

DIST_DIRS := find * -type d -exec

.DEFAULT_GOAL := help

ifndef GOBIN
GOBIN := $(shell echo "$${GOPATH%%:*}/bin")
endif

LINT := $(GOBIN)/golint
GOX := $(GOBIN)/gox
ARCHIVER := $(GOBIN)/archiver
GHR := $(GOBIN)/ghr

$(LINT): ; @go get github.com/golang/lint/golint
$(GOX): ; @go get github.com/mitchellh/gox
$(ARCHIVER): ; @go get github.com/mholt/archiver/cmd/arc
$(GHR): ; @go get github.com/tcnksm/ghr

.PHONY: help
help: ## Show help see: https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: deps
deps: ## Install dependency libraries
	go mod download

.PHONY: tidy
tidy: ## remove unnecessary deps
	go mod tidy

.PHONY: update-deps
update-deps: ## Update dependency libraries
	go get -u

.PHONY: build
build: deps ## Build app for linux arch
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -v -o dist/$(NAME)

.PHONY: run
run: ## Run script
	./bin/go-scraper