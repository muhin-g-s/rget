APP_NAME=rget

all: clean build

build:
	go build -o build/${APP_NAME} ./cmd

run:
	go run cmd/main.go 

clean:
	rm -rf build