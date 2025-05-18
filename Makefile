VERSION := "0.0.1"

test:
	@go test -v -cover ./...

build: generate tidy
	@mkdir -p bin
	@go build -o ./bin/task-exporter ./cmd/task-exporter/main.go

docker:
	@docker build -f ./build/Dockerfile -t task-exporter:${VERSION} .

generate:
	@go generate ./...

tidy:
	@go mod tidy

cleanup:
	@rm -rf ./bin

.PHONY: build