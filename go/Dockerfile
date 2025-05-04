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
FROM cgr.dev/chainguard/static:latest@sha256:2e3db1641bb4fe4e85d2210f4aadb79252e90d5fa745f53a3ffed6a1aab4f73b AS final

LABEL maintainer="rush@ru8733.com"
USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot --chmod=644 /go/src/out/covid-exporter /app

# run binary; use vector form
ENTRYPOINT ["/app"]

EXPOSE 8000
