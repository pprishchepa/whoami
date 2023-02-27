test:
	go test -v ./...
.PHONY: test

bench:
	go test -v -bench=. -benchmem
.PHONY: bench

fmt:
	find . -name \*.go -exec goimports -w {} \;
.PHONY: fmt

mod:
	go mod tidy -v
.PHONY: mod

lint:
	golangci-lint run ./...
.PHONY: lint

finalcheck: fmt mod lint test
.PHONY: finalcheck
