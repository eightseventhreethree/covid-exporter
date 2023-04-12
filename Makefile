CONTAINER_REPO=rushsimonson
CONTAINER_NAME=covid-exporter
CONTAINER_VERSION=0.0.4
CONTAINER_LINUX=$(CONTAINER_NAME)_linux
CONTAINER_ARM=$(CONTAINER_NAME)_arm64

all: stop-build-run

git-tag:
	git tag $(CONTAINER_VERSION)

stop-build-run: stop-container build-container run-container

stop-container:
	docker stop $(CONTAINER_NAME) | true
	docker rm $(CONTAINER_NAME) | true

run-container:
	docker run --detach --publish 8000:8000 --name $(CONTAINER_NAME) -it --restart always $(CONTAINER_REPO)/$(CONTAINER_NAME):$(CONTAINER_VERSION)

build-container:
	docker build --no-cache -t $(CONTAINER_REPO)/$(CONTAINER_NAME):$(CONTAINER_VERSION) .

build-multiarch:
	docker buildx build --push --platform linux/arm64/v8,linux/amd64 --no-cache -t $(CONTAINER_REPO)/$(CONTAINER_NAME):$(CONTAINER_VERSION) .
