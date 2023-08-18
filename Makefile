build:
	@go build -o bin/loyalty-blockchain

run: build
	@./bin/loyalty-blockchain

test:
	@go test -v ./..