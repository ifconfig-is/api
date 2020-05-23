# Build
FROM golang:latest AS build
WORKDIR /geoip2-gql
COPY . .
ENV GO111MODULE=on CGO_ENABLED=0
RUN go build .

# Run
FROM scratch
COPY --from=build \
	/geoip2-gql/geoip2-gql \
	/geoip2-gql

ENV GEOIP2GQL_PORT=3000
ENV MAXMIND_PATH=/maxmind

ENTRYPOINT ["/geoip2-gql"]
