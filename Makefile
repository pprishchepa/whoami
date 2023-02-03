.PHONY: fmt
fmt:
	find . -name \*.go -not -path \*/wire_gen.go -exec goimports -w {} \;

.PHONY: mod
mod:
	go mod tidy -v

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: finalcheck
finalcheck: fmt mod lint
