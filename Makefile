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

L1POOL_ABI_ARTIFACT := bindings/abi/L1PoolManager.json
L2POOL_ABI_ARTIFACT := bindings/abi/L2PoolManager.json
STRATEGYBASE_ABI_ARTIFACT := bindings/abi/StrategyBase.json
STRATEGYMANAGER_ABI_ARTIFACT := bindings/abi/StrategyManager.json
STAKINGMANAGER_ABI_ARTIFACT := bindings/abi/StakingManager.json
DETH_ABI_ARTIFACT := bindings/abi/DETH.json

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
	| jq -r .bytecode.object > $(temp)

	cat $(L1POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_l1_pool_manager.go \
	--type L1PoolManager \
	--bin $(temp)

	rm $(temp)

binding-l2p:
	$(eval temp := $(shell mktemp))

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(L2POOL_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_l2_pool_manager.go \
	--type L2PoolManager \
	--bin $(temp)

	rm $(temp)

binding-strategyM:
	$(eval temp := $(shell mktemp))

	cat $(STRATEGYMANAGER_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(STRATEGYMANAGER_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_strategy_manager.go \
	--type StrategyManager \
	--bin $(temp)

	rm $(temp)

binding-strategyB:
	$(eval temp := $(shell mktemp))

	cat $(STRATEGYBASE_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(STRATEGYBASE_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_strategy_base.go \
	--type StrategyBase \
	--bin $(temp)

	rm $(temp)

binding-stakingM:
	$(eval temp := $(shell mktemp))

	cat $(STAKINGMANAGER_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(STAKINGMANAGER_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_staking_manager.go \
	--type StakingManager \
	--bin $(temp)

	rm $(temp)

binding-dETH:
	$(eval temp := $(shell mktemp))

	cat $(DETH_ABI_ARTIFACT) \
	| jq -r .bytecode.object > $(temp)

	cat $(DETH_ABI_ARTIFACT) \
	| jq .abi \
	| abigen --pkg bindings \
	--abi - \
	--out bindings/bvm_deth.go \
	--type DETH \
	--bin $(temp)

	rm $(temp)

.PHONY: all build test clean run deps binding-l1p binding-l2p binding-strategyM binding-strategyB binding-stakingM binding-dETH

