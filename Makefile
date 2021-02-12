CONTAINER_NAME=covid-exporter
CONTAINER_LINUX=$(CONTAINER_NAME)_linux
CONTAINER_ARM=$(CONTAINER_NAME)_arm64
TAG_NAME:=$(shell git describe --abbrev=0 --tags)

ifeq ($(origin TAG_NAME),"")
TAG_NAME:="latest"
endif

all: build-container run-container

stop-build-run: stop-container build-container run-container

stop-container:
	docker stop $(CONTAINER_NAME) | true
	docker rm $(CONTAINER_NAME) | true

run-container:
	docker run --detach --publish 8000:8000 --name $(CONTAINER_NAME) -it --restart always $(CONTAINER_NAME):$(TAG_NAME)

build-container:
	docker build --no-cache -t $(CONTAINER_NAME):$(TAG_NAME) .
