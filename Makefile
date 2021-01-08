GO=go
GO_BUILD=$(GO) build
GO_CLEAN=$(GO) clean
GO_TEST=$(GO) test
GO_MOD=$(GO) mod
WIRE=~/go/bin/wire

APP_DIR=app
BIN_DIR=bin
CMD_DIR=app/cmd
SCRIPTS_DIR=scripts
BIN_DIR_RELATIVE_CMD=../../$(BIN_DIR)

MAIN_TARGET = server
ALL_TARGETS = $(MAIN_TARGET) foo bar
WIRE_TARGETS = $(patsubst %,wire-%, $(ALL_TARGETS))
BUILD_TARGETS = $(patsubst %,build-%, $(ALL_TARGETS))
RUN_TARGETS = $(patsubst %,run-%, $(ALL_TARGETS))

HOTH=hoth-operacoes
HOTH_BIN_OUT=../../$(HOTH)

all: help

help: ## Display help message
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) \
	| sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\\x1b[36m\1\\x1b[m:\2/' \
	| echo "$$(cat -)\n$(WIRE_TARGETS): Updates wire_gen.go for specific target " \
	| echo "$$(cat -)\n$(BUILD_TARGETS):  " \
	| echo "$$(cat -)\n$(RUN_TARGETS):  " \
	| sort -r  \
	| column -c2 -t -s :)"

$(ALL_TARGETS):
	@echo $@
	
wire: wire-$(MAIN_TARGET) ## Updates wire_gen.go for main target

wire-all: $(WIRE_TARGETS) ## Updates wire_gen.go for all targets

$(WIRE_TARGETS): wire-%: %
	cd $(CMD_DIR)/$< && $(WIRE)

tidy: ## Updates go.mod and go.sum
	cd $(APP_DIR) && $(GO_MOD) tidy

test: ## Test the app
	cd $(APP_DIR) && $(GO_TEST) ./...

build: build-$(MAIN_TARGET) ## Builds a local version for main target

build-all: $(BUILD_TARGETS) ## Builds a local version for all targets

$(BUILD_TARGETS): build-%: %
	cd $(CMD_DIR)/$< && $(GO_BUILD) -o ../$(BIN_DIR_RELATIVE_CMD)
	cp -R res $(BIN_DIR)

clean: ## Remove binaries
	rm -f $(BIN_DIR)/*

refresh: wire wire-all  test build clean ## Make wire tidy test build and clean at once

run: run-$(MAIN_TARGET) ## Run a local version

$(RUN_TARGETS): run-%: % build-%
	$(BIN_DIR)/$<

.PHONY: help wire wire-all $(WIRE_TARGETS) clean wire tidy build run test refresh benchmarks $(TARGETS)

# not ready


integration-test: ## Run integration tests
	@echo TODO

benchmark: ## Run benchmark
	@echo TODO

