default: vet fmt

check: vet test

.PHONY: test
test:
	@go test -v ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: fmt
fmt:
	@go fix ./...
	@go fmt ./...
