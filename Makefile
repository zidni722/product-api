PROJECT_NAME=product-api

hello:
	echo "Hello"

build:
	go clean
	go build

run:
	./${PROJECT_NAME}

build_and_run: build run