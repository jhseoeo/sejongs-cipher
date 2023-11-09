BINARY_NAME=app

build:
	go build -o ./build/$(BINARY_NAME) -v

run: build
	./build/$(BINARY_NAME)

clean:
	go clean
	rm -f ./build/$(BINARY_NAME)