.PHONY: run
run:
	./cmd/run-api 

.PHONY: lint
lint:
	@golangci-lint run -v

.PHONY: validate
validate:
	@go test -v ./...
	@golangci-lint run -v
	@go build -v ./...

.PHONY: schema
schema:
	@migrate create -ext sql -dir src/migration/scripts -seq $(name)

.PHONY: migrate
migrate:
	@go run src/migration/main.go