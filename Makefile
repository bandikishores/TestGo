GOPRIVATE="github.org/bandikishores,bandikishores.com/bandikishorescom,TestGo,bandi.com/TestGo"
GOPROXY = direct
PACKAGE  = common
DATE     = $(shell date +%s)
BIN      = $(GOPATH)/bin
BASE     = $(PWD)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

PATCH_VERSION ?= $(shell if [ -z "$(BITBUCKET_BUILD_NUMBER)" ] && [ "$(BITBUCKET_BUILD_NUMBER)xxx" == "xxx" ]; then echo "dev"; else echo $(BITBUCKET_BUILD_NUMBER); fi)
# This needs to be manually updated according to semver rules
VERSION ?= "v0.1.$(PATCH_VERSION)"

GO       = go
GODOC    = godoc
GOFMT    = gofmt
DOCKER   = docker
TIMEOUT  = 15

V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")

.PHONY: all
all: fmt gen lint build

# Tools
GOMOD = $(GO) mod

GOLINT = $(BIN)/golint
$(BIN)/golint: | $(BASE) ; $(info $(M) installing golint…)
	$Q $(GO) get -u golang.org/x/lint/golint

PROTOC_GRPC_GATEWAY = $(BIN)/protoc-gen-grpc-gateway
$(BIN)/protoc-gen-grpc-gateway: | $(BASE) ; $(info $(M) installing protoc-gen-grpc-gateway…)
	$Q $(GO) get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

PROTOC_SWAGGER = $(BIN)/protoc-gen-swagger
$(BIN)/protoc-gen-swagger: | $(BASE) ; $(info $(M) installing protoc-gen-swagger…)
	$Q $(GO) get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

PROTOC_GO = $(BIN)/protoc-gen-go
$(BIN)/protoc-gen-go: | $(BASE) ; $(info $(M) installing protoc-gen-go…)
	$Q $(GO) get -u github.com/golang/protobuf/protoc-gen-go

PROTOC_GEN_GOGO = $(BIN)/protoc-gen-gogo
$(BIN)/protoc-gen-gogo: | $(BASE) ; $(info $(M) installing protoc-gen-gogo…)
	$Q $(GO) get -u github.com/gogo/protobuf/protoc-gen-gogo

GOOGLE_APIS = $(BIN)/google-apis
$(BIN)/google-apis: | $(BASE) ; $(info $(M) installing googleapis…)
	$Q $(GO) get -u github.com/googleapis/googleapis@v0.0.0-20200303215514-541b1ded4aba

# Dependency management
.PHONY: download
download:
	$Q cd $(BASE) && $(GO) mod download; $(info $(M) retrieving dependencies…)

.PHONY: install-dependencies 
install-dependencies: $(BIN)/golint $(BIN)/protoc-gen-grpc-gateway $(BIN)/protoc-gen-swagger $(BIN)/protoc-gen-go $(PROTOC_GEN_GOGO) $(GOOGLE_APIS)  ## Install dependent go tools

# Tests
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
test: fmt lint | $(BASE) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q cd $(BASE) && $(GO) test -v -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

.PHONY: bench
bench: ; $(info $(M) running $(NAME:%=% )benchmarks…) @ ## Run benchmarks
	$(GO) test ./... -bench=. -run=^$$

.PHONY: lint
lint: download | $(BASE) $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: gen
gen:
	$Q $(GO) mod download
	$Q cd $(BASE) && $(GO) generate ./pkg/data; $(info $(M) generating api...)

.PHONY: build
build: gen  ## Build service
	$Q cd $(BASE) && $(GO) build ./... ; $(info $(M) attempting build…)
.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin

# Create a new tag if the current branch is being merged into master
.PHONY: master-tag
master-tag:
ifeq ($(BITBUCKET_BRANCH), master)
	git tag $(VERSION)
	git push origin --tags
endif

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'