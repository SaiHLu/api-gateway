clean:
	@rm -rf bin

build: clean
	@go build -o bin/gateway cmd/main.go

run: build
	@./bin/gateway

generate-swagger:
	@swag init -g main.go -d ./cmd,./internal -o ./docs --parseFuncBody --parseInternal --parseDependency

# -o docs --parseDependency --parseInternal
.PHONY: build run