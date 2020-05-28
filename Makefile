export CGO_ENABLED=0
export GO111MODULE=on

IMAGE_PROCESSING_CONTROLLER_VERSION := latest

.PHONY: build



coverage: # @HELP generate unit test coverage data
coverage: build linters

build: # @HELP build the source code
build: deps
	GOOS=linux GOARCH=amd64 go build -o build/_output/image-processing-controller ./cmd/image-processing-controller

deps: # @HELP ensure that the required dependencies are in place
	#go build -v ./...
	#bash -c "diff -u <(echo -n) <(git diff go.mod)"
	#bash -c "diff -u <(echo -n) <(git diff go.sum)"

linters: # @HELP examines Go source code and reports coding problems
	GOGC=50 golangci-lint run


images: # @HELP build image processing controller Docker image
image: build
	docker build . -f build/image-processing-controller/Dockerfile -t  image-processing-controller:${IMAGE_PROCESSING_CONTROLLER_VERSION}

kind: image
	kind load docker-image image-processing-controller:${IMAGE_PROCESSING_CONTROLLER_VERSION}
	

all: test


clean: # @HELP remove all the build artifacts
	rm -rf ./build/_output ./vendor

help:
	@grep -E '^.*: *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '
