APP_NAME=rget

all: clean build

build:
	go build -o build/${APP_NAME} ./cmd

run:
	go run cmd/main.go $(filter-out $@,$(MAKECMDGOALS))

test:
	go test ./... -v

clean:
	rm -rf build

%:
	@: