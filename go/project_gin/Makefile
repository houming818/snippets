# Set shell to bash
SHELL := /bin/bash

# TARGETS defines the app to build.
TARGETS := app

# REGISTRY defines the registry to push images.
REGISTRY := ccr.ccs.tencentyun.com/grepcode

# IMAGE_TAG defines the image tag.
IMAGE_TAG := latest

# IMAGE defines the full image name.
IMAGE := $(REGISTRY)/$(TARGETS):$(IMAGE_TAG)

GOARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))

# Project main package location (can be multiple ones).
CMD_DIR := ./cmd

# Root of this project.
ROOT := 

# Whether enable cgo
CGO_ENABLED := 1

# define some go build vars
# define some go build vars
# Set DEBUGGER=1 to build debug symbols
#LDFLAGS = $(if $(DEBUGGER),,-s -w) $(shell ROOT=$(ROOT) ./hacks/version.sh)

# Project output directory.
OUTPUT_DIR := ./target

# Build direcotory.
BUILD_DIR := ./build

# Git commit sha.
COMMIT := $(shell git rev-parse --short HEAD)

# Git tag describe.
TAG = $(shell git rev-parse --short HEAD)

# Current version of the project.
VERSION ?= $(shell cat version)

# Project name
PROJECT_NAME := app

# Golang standard bin directory.
BIN_DIR := $(GOPATH)/bin

MKDIR = mkdir

# All targets.
.PHONY: build lint test image run

lint: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) run

install:
	GO111MODULE=on
	@go mod tidy
	@go mod download

test:
	@go mod vendor
	@go test -v $$(go list ./... | grep -E -v '/vendor|/test|/mock') -coverprofile=coverage.out
	@go tool cover -func=coverage.out | tail -n 1 | awk '{ print "Total coverage: " $$3 }'
	@go tool cover -html=coverage.out

run: install
	go run cmd/$(PROJECT_NAME)/main.go

compile:
	CGO_ENABLED=0 GO111MODULE=on
	go mod vendor
	go build -mod=vendor -v -o $(OUTPUT_DIR)/$(PROJECT_NAME) -ldflags '$(LDFLAGS)' ./cmd/main.go

#golangci_lint:
#	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(BIN_DIR) v1.30.0

# 使用docker进行交叉编译
build:
	docker build --rm --no-cache -t $(IMAGE) -f build/Dockerfile .

push:
	docker push $(IMAGE) 

.PHONY: clean
clean:
	-rm ${OUTPUT_DIR}/*
