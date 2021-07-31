# Build
FROM golang:1.16.6-buster AS build
ARG DEBIAN_FRONTEND=noninteractive
WORKDIR /
COPY . .

# Build api
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build .

# Run
FROM scratch
COPY --from=build \
	/api /api

ENTRYPOINT ["/api"]
