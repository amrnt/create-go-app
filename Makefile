# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

PACKAGE = github.com/amrnt/create-go-app
COMMIT_HASH = `git rev-parse --short HEAD 2>/dev/null`
BUILD_DATE = `date +%FT%T%z`
LDFLAGS = -ldflags "-X ${PACKAGE}/pkg/config.CommitHash=${COMMIT_HASH} -X ${PACKAGE}/pkg/config.BuildDate=${BUILD_DATE}"

.PHONY: vendor help
.DEFAULT_GOAL := help

vendor: ## Install glide and install the dependencies
	go get github.com/Masterminds/glide
	glide install --strip-vendor

LOCAL_PACKAGES = `go list ./... | grep -v /vendor/`

test: ## Test the code excluding vendor dir
	go test -v ${LOCAL_PACKAGES}

run: ## Run run run
	go run main.go ${CMD}

build-only: ## Build the binary in the current dir without running `vendor`
	go build ${LDFLAGS} ${PACKAGE}

build: vendor ## Build the binary in the current dir
	go build ${LDFLAGS} ${PACKAGE}

build-race: vendor ## Build the binary with race detector enabled
	go build -race ${LDFLAGS} ${PACKAGE}

install: vendor ## Install the binary
	go install ${LDFLAGS} ${PACKAGE}

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
