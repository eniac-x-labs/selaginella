# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Main package path
MAIN_PATH=./cmd

# Binary names
BINARY_NAME=selaginella

L1POOL_ABI_ARTIFACT := bindings/abi/L1Pool.json
L2POOL_ABI_ARTIFACT := bindings/abi/L2Pool.json

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET) -v ./...

binding-l1p:
	$(eval temp := $(shell mktemp))

	cat $(L1POOL_ABI_ARTIFACT) \
	| jq -r .bytecode > $(temp)

	cat $(L1POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_l1_pool.go \
	--type L1Pool \
	--bin $(temp)

	rm $(temp)

binding-l2p:
	$(eval temp := $(shell mktemp))

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq -r .bytecode > $(temp)

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_l2_pool.go \
	--type L2Pool \
	--bin $(temp)

	rm $(temp)

.PHONY: all build test clean run deps binding-l1p binding-l2p

