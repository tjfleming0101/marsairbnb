run: build
	@./bin/marsairbnb

build:
	@go build -o bin/marsairbnb ./cmd/web/ .