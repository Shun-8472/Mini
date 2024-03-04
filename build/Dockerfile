# syntax = docker/dockerfile:experimental
FROM golang:latest

# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# final stage
FROM centurylink/ca-certs
COPY --from=build-env /src/app /
COPY --from=build-env /src/config /config
ENTRYPOINT ["/app"]