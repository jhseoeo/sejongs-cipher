BINARY_NAME=app

prepare: clean
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init --parseDependency -g routes.go -o ./internal/interfaces/docs -d ./internal/interfaces/routes

run: prepare
	go run ./...

build: prepare
	go build -o ./build/ -v ./...

clean:
	go clean
	rm -f ./build/$(BINARY_NAME)