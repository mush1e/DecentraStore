build:
	@go build -o bin/dcs

run: build
	@./bin/dcs

test:
	@go test ./... -v