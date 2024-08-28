# argument for Go version
ARG GO_VERSION=1.21.6

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-bookworm AS build

RUN apt-get update && apt-get -y install make
WORKDIR $GOPATH/src
ADD . $GOPATH/src
RUN make buildx-build


# STAGE 2: build the container to run
FROM cgr.dev/chainguard/static:latest@sha256:791657dd88dea8c1f9d3779815429f9c681a9a2778fc66dac3fbf550e1f1d9c8 AS final

LABEL maintainer="rush@ru8733.com"
USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot /go/src/out/covid-exporter /app

# run binary; use vector form
ENTRYPOINT ["/app"]

EXPOSE 8000
