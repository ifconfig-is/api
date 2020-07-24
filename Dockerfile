# Build
FROM golang:latest AS build
WORKDIR /api
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build .

# Run
FROM scratch
COPY --from=build \
	/api/api \
	/api

ENV GEOIP2GQL_PORT=3000
ENV MAXMIND_PATH=/maxmind

ENTRYPOINT ["/api"]
