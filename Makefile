.PHONY: run
run:
	@go run main.go


.PHONY: lint
lint:
	@go mod tidy
	@gofmt -e -s -w .
	@goimports -v -w .
	@golint .
	@go vet