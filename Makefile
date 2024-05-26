# Copyright 2024 Gustavo Salomão
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Project parameters
NAME          := terracov
BUILD_DIR     := dist
COVERAGE_DIR  := coverage
TEST_TIMEOUT  := 10s
NO_COLORS     ?= false
GOARCH        ?= $(shell go env GOARCH)
GOOS          ?= $(shell go env GOOS)
VERSION       ?= $(shell git describe --tags --always --dirty | sed -e 's/^v//')
REVISION      ?= $(shell git rev-parse --short HEAD)
BUILD_TIME    ?= $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS       := "\
	-X 'github.com/gsalomao/terracov/internal/build.version=${VERSION}' \
	-X 'github.com/gsalomao/terracov/internal/build.revision=${REVISION}' \
	-X 'github.com/gsalomao/terracov/internal/build.buildTime=${BUILD_TIME}' \
	"

# Colors
ifeq ($(NO_COLORS),false)
	GREEN     := $(shell tput -Txterm setaf 2)
	YELLOW    := $(shell tput -Txterm setaf 3)
	WHITE     := $(shell tput -Txterm setaf 7)
	CYAN      := $(shell tput -Txterm setaf 6)
	RESET     := $(shell tput -Txterm sgr0)
	BOLD      := \033[0;1m
endif

.PHONY: all
all: help

## Project
.PHONY: init
init: ## Initialize repository
	@pre-commit install --hook-type pre-commit
	@pre-commit install --hook-type commit-msg

.PHONY: fmt
fmt: ## Format source code
	@go fmt ./...

## Build
.PHONY: build
build: ## Build application
	@mkdir -p ${BUILD_DIR}
	@go build -o ${BUILD_DIR}/$(NAME) -ldflags ${LDFLAGS} cmd/$(NAME)/main.go

.PHONY: clean
clean: ## Clean build folder
	@go clean
	@rm -rf ${BUILD_DIR}

## Test
.PHONY: test
test: ## Run unit tests
	@gotestsum --format testname --packages ./... -- -timeout ${TEST_TIMEOUT} -race

.PHONY: test-ci
test-ci: ## Run unit tests on CI
	@go test ./... -timeout ${TEST_TIMEOUT} -race

.PHONY: coverage
coverage: ## Run unit tests with coverage report
	@rm -rf ${COVERAGE_DIR}
	@mkdir -p ${COVERAGE_DIR}
	@go test -timeout ${TEST_TIMEOUT} -cover -covermode=atomic -race \
			-coverprofile=${COVERAGE_DIR}/coverage.out ./...
	@cat ${COVERAGE_DIR}/coverage.out | \
			grep -v "mocks" | grep -v "main.go" > \
			${COVERAGE_DIR}/coverage.final.out
	@mv ${COVERAGE_DIR}/coverage.final.out ${COVERAGE_DIR}/coverage.out
	@go tool cover -func ${COVERAGE_DIR}/coverage.out

.PHONY: coverage-html
coverage-html: coverage ## Open the coverage report in the browser
	@go tool cover -html ${COVERAGE_DIR}/coverage.out

## Inspect
.PHONY: vet
vet: ## Examine source code
	@go vet ./...

.PHONY: lint
lint: ## Lint source code
	@golint  -set_exit_status $(go list ./...)
	@golangci-lint run $(go list ./...)

.PHONY: complexity
complexity: ## Calculates cyclomatic complexity
	@gocyclo -top 10 .
	@gocyclo -over 15 -avg .

.PHONY: security
security: ## Run security checks
	@govulncheck ./...
	@gosec -quiet -exclude-dir testdata ./...

## Help
.PHONY: help
help: ## Show this help
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@awk 'BEGIN {FS = ":.*?## "} { \
			if (/^[a-zA-Z0-9\/_-]+:.*?##.*$$/) { \
					printf "  ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
					else if (/^## .*$$/) { \
							printf "${CYAN}%s:${RESET}\n", substr($$1,4)\
					} \
			}' $(MAKEFILE_LIST)
