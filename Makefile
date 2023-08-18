build:
	@go build -o bin/loyalty-blockchain

run: build
	@./bin/loyalty-blockchain

test:
	@go test -v ./..

compile_sol:
	@sh scripts/compile_sol.sh

sol: compile_sol
	@abigen --abi=./blockchain/abigenBindings/abi/LoyaltyProgramme.abi --bin=./blockchain/abigenBindings/bin/LoyaltyProgramme.bin --pkg=api --out=./application/api/LoyaltyProgramme.go