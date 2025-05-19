VERSION := "0.0.1"
NAME    := "task-exporter"

test:
	@go test -v -cover ./...

build:
	@mkdir -p bin
	@go build -o ./bin/task-exporter ./cmd/task-exporter/main.go

docker_build: test
	@docker build -f ./build/Dockerfile -t ${NAME}:${VERSION} .

docker_run:
	@docker run --name ${NAME} -p 127.0.0.1:8080:8080 ${NAME}:0.0.1

generate:
	@go generate ./...

tidy:
	@go mod tidy

cleanup:
	@rm -rf ./bin

.PHONY: build