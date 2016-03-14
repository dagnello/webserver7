all: 
		clean build test

clean: 
		rm -rf ${GOBIN}/webserver7

build: 
		godep go install

test: 
		godep go test ./...
