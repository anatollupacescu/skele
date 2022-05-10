PROJECT:=$(shell go list -m)

GOOS?=linux
GOARCH?=amd64

RELEASE?=0.0.1
COMMIT := git-$(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: format test clean

format:
	@goimports -w -local $(PROJECT) cmd/skele
	@gci -w -local $(PROJECT) .
	@gofmt -s -w cmd/skele

test:
	@go test -count 1 -race -cover ./cmd/skele/...

clean:
	rm -rf cmd/skele/list cmd/skele/item_test.go cmd/skele/doc.go