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
NO_COLORS     ?= false
GOARCH        ?= $(shell go env GOARCH)
GOOS          ?= $(shell go env GOOS)
VERSION       ?= $(shell git describe --tags --always --dirty | sed -e 's/^v//')
REVISION      ?= $(shell git rev-parse --short HEAD)

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
