# argument for Go version
ARG GO_VERSION=1.24.3

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-bookworm AS build

RUN apt-get update && apt-get -y install make
WORKDIR $GOPATH/src
ADD . $GOPATH/src
RUN make buildx-build


# STAGE 2: build the container to run
FROM cgr.dev/chainguard/static:latest@sha256:797e62f43d04d792e9f930913e7d9f5a63e92bd19ca5e7e5139a692decee2dbc AS final

LABEL maintainer="rush@ru8733.com"
USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot /go/src/out/covid-exporter /app

# run binary; use vector form
ENTRYPOINT ["/app"]

EXPOSE 8000
