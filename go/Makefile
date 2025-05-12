# Go parameters
CONTAINER_REPO=rushsimonson
CONTAINER_NAME=covid-exporter
GOCMD=go
BINARY_NAME=$(CONTAINER_NAME)
BINARY_LINUX=$(BINARY_NAME)_linux
BINARY_BUILDX=$(BINARY_NAME)
BUILD_BASE=$(GOCMD) build -o out/
CONTAINER_VERSION=go-0.1.0

all: test build
build:
	$(BUILD_BASE)$(BINARY_NAME) -v

test:
	$(GOCMD) test -v ./...

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_LINUX)

vendor:
	$(GOCMD) mod vendor

run:
	$(GOCMD) run main.go

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(BUILD_BASE)$(BINARY_LINUX) -v

# Buildx make cmd
buildx-build:
	CGO_ENABLED=0 $(BUILD_BASE)$(BINARY_BUILDX) -v

build-multiarch:
	docker buildx build --push --platform linux/arm64/v8,linux/amd64 --no-cache -t $(CONTAINER_REPO)/$(CONTAINER_NAME):$(CONTAINER_VERSION) .

build-container:
	docker build --no-cache -t $(CONTAINER_NAME):$(CONTAINER_VERSION) .
