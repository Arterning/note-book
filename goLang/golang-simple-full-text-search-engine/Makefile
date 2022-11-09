all: help

.PHONY : help
help : Makefile
	@sed -n 's/^##//p' $< | awk 'BEGIN {FS = ":"}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

TOOLS_MOD_DIR := ./tools
TOOLS_DIR := $(abspath ./.tools)
$(TOOLS_DIR)/golangci-lint: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	@echo BUILD golangci-lint
	@cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(TOOLS_DIR)/mockery: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	@echo BUILD mockery
	@cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/mockery github.com/vektra/mockery/v2

$(TOOLS_DIR)/cobra: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	@echo BUILD cobra
	@cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/cobra github.com/spf13/cobra/cobra

$(TOOLS_DIR)/godoc: $(TOOLS_MOD_DIR)/go.mod $(TOOLS_MOD_DIR)/go.sum $(TOOLS_MOD_DIR)/tools.go
	@echo BUILD godoc
	@cd $(TOOLS_MOD_DIR) && \
	go build -o $(TOOLS_DIR)/godoc golang.org/x/tools/cmd/godoc

## tools: Build all tools
tools: $(TOOLS_DIR)/mockery $(TOOLS_DIR)/golangci-lint $(TOOLS_DIR)/cobra

## lint: Run golangci-lint
.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint gen
	@echo LINT
	@$(TOOLS_DIR)/golangci-lint run -c .github/linters/.golangci.yaml --out-format colored-line-number
	@printf "LINT... \033[0;32m [OK] \033[0m\n"

## test: Run small test
.PHONY: test
test: gen
	@echo SMALL TEST
	@go test -test.short ./...
	@printf "SMALL TEST... \033[0;32m [OK] \033[0m\n"

## test/medium: Run medium test
test/medium: gen
	@echo MEDIUM TEST
	@rm -rf test/report
	@go test ./...
	@printf "MEDIUM TEST... \033[0;32m [OK] \033[0m\n"

## test/coverage: Run small test and generate coverage report
.PHONY: test/coverage
test/coverage:
	@rm -rf test/apitest
	@go test -test.short ./... -coverprofile=coverage.txt -covermode=atomic
	@go tool cover -html=coverage.txt -o coverage.html

## build: Build all binary
BIN_DIR := $(abspath ./bin)
BUILD_TARGETS=build/api \
			  build/cli
build: $(BUILD_TARGETS)

.PHONY: $(BIN_DIR)
$(BIN_DIR):
	@mkdir -p $@

## build/api: Build api server binary
.PHONY: $(BUILD_TARGETS)
build/api: $(BIN_DIR)
	@echo BUILD api
	@go build -o ./bin/api ./api

## build/cli: Build cli binary
.PHONY: $(BUILD_TARGETS)
build/cli: $(BIN_DIR)
	@echo BUILD cli
	@go build -o ./bin/cli ./cli

## gen: Run all code generator
GEN_TARGETS=gen/mock
.PHONY: gen
gen: $(GEN_TARGETS)

## gen/mock: Run mock generator
.PHONY: $(GEN_TARGETS)
gen/mock: $(TOOLS_DIR)/mockery
	@echo GENERATE mocks
	@go generate ./...

## godoc: View godoc
PKG_NAME:=$(shell cat go.mod | grep module | cut -d' ' -f2)
.PHONY: godoc
godoc: $(TOOLS_DIR)/godoc
	@echo "Open http://localhost:6060/pkg/$(PKG_NAME) on browser."
	$(TOOLS_DIR)/godoc -http localhost:6060

## download: Download englush wikipedia dump
.PHONY: download
download:
	@echo DOWNLOAD
	@rm -rf ./test/testdata
	@mkdir -p ./test/testdata
	@wget -O ./test/testdata/enwiki-latest-abstract1.xml.gz wget https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz

# git hooks
.PHONY: pre-push
pre-push: test lint

.PHONY: lint-ci
lint-ci: $(TOOLS_DIR)/golangci-lint gen
	@echo LINT
	@$(TOOLS_DIR)/golangci-lint run -c .github/linters/.golangci.yaml
	@printf "LINT... \033[0;32m [OK] \033[0m\n"
