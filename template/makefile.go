package template

// Makefile template for Makefile
const Makefile = `version?=latest
commit_sha=$(shell git rev-parse --short HEAD)
img={{.DockerImg}}:$(version)
imgdev={{.DockerImg}}-dev:$(version)
uid=$(shell id -u $$USER)
gid=$(shell id -g $$USER)
dockerbuilduser=--build-arg USER_ID=$(uid) --build-arg GROUP_ID=$(gid)
wd=$(shell pwd)
appvol=$(wd):/app
rundev=docker run -it --rm -v $(appvol) $(imgdev)
runbuild=docker run --rm -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 -v $(appvol) $(imgdev)
ldflags="-w -s -X {{.Module}}/pkg/version.Semver=$(version) -X {{.Module}}/pkg/version.GitSHA=$(commit_sha)"
cov=coverage.out
covhtml=coverage.html

.PHONY: imagedev
imagedev: ##@development build image docker dev
	docker build . --target dev $(dockerbuilduser) -t $(imgdev)

.PHONY: githooks
githooks: ##@development install git hooks
	@echo "copying git hooks"
	@mkdir -p .git/hooks
	@cp hack/githooks/pre-commit .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit
	@echo "git hooks copied"

.PHONY: shell
shell: imagedev ##@development open shell in the development image
	$(rundev) bash

.PHONY: release
release: publish ##@production create a git tag and build and publish a docker image
	git tag -a $(version) -m "Generated release "$(version)
	git push origin $(version)

.PHONY: lint
lint: imagedev ##@lint static analysis code base
	$(rundev) golangci-lint run --enable-all

.PHONY: test
test: imagedev ##@test run all tests
	$(rundev) go test ./... -race -covermode=atomic -coverprofile=$(cov) -timeout 30s -v

.PHONY: coverage
coverage: test ##@test coverage all tests package
	$(rundev) go tool cover -html=$(cov) -o=$(covhtml)
	xdg-open coverage.html

.PHONY: build
build: imagedev ##@development Build binary
	$(runbuild) go build -v -ldflags $(ldflags) -o ./cmd/{{.Project}}/{{.Project}} ./cmd/{{.Project}}

.PHONY: image
image: ##@production Build docker image
	docker build . --build-arg LDFLAGS=$(ldflags) -t $(img)

.PHONY: publish
publish: image ##@production Build and publish docker image
	docker push $(img)

.PHONY: modtidy
modtidy: imagedev ##@development add missing and remove unused modules
	$(rundev) go mod tidy

.PHONY: fmt
fmt: imagedev ##@development fmt project
	$(rundev) gofmt -w -s -l .

.PHONY: run
run: image ##@development execute docker image
	docker run --rm -p 8080:8080 $(img)

.DEFAULT_GOAL := help
#COLORS
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_FUN = \
	%help; \
	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	print "usage: make [target]\n\n"; \
	for (sort keys %help) { \
	print "${WHITE}$$_:${RESET}\n"; \
	for (@{$$help{$$_}}) { \
	$$sep = " " x (32 - length $$_->[0]); \
	print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; }

.PHONY: help
help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)
`
