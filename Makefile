GOOS ?= windows
GOARCH ?= amd64
CGO_ENABLED ?= 1 
COMMANDS=$(wildcard ${SRCDIR}/cmd/*)

ifeq ($(NOGIT),1)
  GIT_SUMMARY ?= Unknown
  GIT_BRANCH ?= Unknown
  GIT_MERGE ?= Unknown
else
  GIT_SUMMARY := $(shell ${GITFLAGS} git describe --tags --dirty --always)
  GIT_BRANCH := $(shell ${GITFLAGS} git symbolic-ref -q --short HEAD)
  GIT_MERGE := $(shell ${GITFLAGS} git rev-list --count --merges master)
endif

.PHONY: setup
setup:
	@echo "Setting up project"
	@echo
	@echo "Download dependencies"
	go mod download

	@echo
	@echo "Generating"
	go generate -v ./...

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	@echo "Building application"
	@echo
	go build

.PHONY: serve
serve:
	@docker run --rm -it -p 8000:8000 -v ${PWD}/www:/docs squidfunk/mkdocs-material

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo