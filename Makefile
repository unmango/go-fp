GO        ?= go
GINKGO    ?= ginkgo
GOMOD2NIX ?= gomod2nix
NIX       ?= nix

GO_SRC := $(shell find . -name '*.go')

ifeq ($(CI),)
TEST_FLAGS := --label-filter !E2E
else
TEST_FLAGS := --github-output --race --trace --coverprofile=cover.profile
endif

build:
	$(NIX) build --no-substitute

test:
	$(GINKGO) run -r ${TEST_FLAGS}

check:
	$(NIX) flake check

update:
	$(NIX) flake update

tidy: go.sum nix/gomod2nix.toml

clean:
	find . -name report.json -delete

go.sum: go.mod ${GO_SRC}
	go mod tidy
	@touch $@

nix/gomod2nix.toml: go.mod go.sum
	$(GOMOD2NIX) generate --outdir nix

%_suite_test.go:
	cd ${@D} && $(GINKGO) bootstrap

%_test.go:
	cd ${@D} && $(GINKGO) generate $(notdir $*)
