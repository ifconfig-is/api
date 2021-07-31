FROM scratch

WORKDIR /
COPY ./api ./api

ENTRYPOINT ["/api"]
