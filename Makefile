build:
	@go build -o bin/ScyllaTaskify cmd/todo/main.go

run: build
	@./bin/ScyllaTaskify