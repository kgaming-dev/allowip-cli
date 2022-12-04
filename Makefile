BINARY_NAME=allowip
BUILD_DIR=build/amd64-linux

build:
 GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME} main.go

run:
 ./{$BUILD_DIR}/${BINARY_NAME}

build_and_run: build run

clean:
 go clean
 rm -rf ${BUILD_DIR}
