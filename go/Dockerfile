# argument for Go version
ARG GO_VERSION=1.21.6

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-bookworm AS build

RUN apt-get update && apt-get --no-install-recommends -y install make && \
  rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*
WORKDIR $GOPATH/src
ADD . $GOPATH/src
RUN make buildx-build


# STAGE 2: build the container to run
FROM cgr.dev/chainguard/static:latest@sha256:bda736ff75ec4b76ecd120d6ccf0b9428419cf4568a5d786a7d77ba23a9f6992 AS final

LABEL maintainer="rush@ru8733.com"
USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot --chmod=644 /go/src/out/covid-exporter /app

# run binary; use vector form
ENTRYPOINT ["/app"]

EXPOSE 8000
