# argument for Go version
ARG GO_VERSION=1.20.3

# STAGE 1: building the executable
FROM golang:${GO_VERSION}-bullseye AS build

RUN apt update && apt install make
WORKDIR $GOPATH/src
ADD . $GOPATH/src
RUN make buildx-build


# STAGE 2: build the container to run
FROM gcr.io/distroless/static AS final

LABEL maintainer="rush@ru8733.com"
USER nonroot:nonroot

# copy compiled app
COPY --from=build --chown=nonroot:nonroot /go/src/out/covid-exporter /app

# run binary; use vector form
ENTRYPOINT ["/app"]

EXPOSE 8000