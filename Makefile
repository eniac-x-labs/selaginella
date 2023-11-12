GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

selaginella:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/selaginella

clean:
	rm selaginella

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	selaginella \
	clean \
	test \
	lint
