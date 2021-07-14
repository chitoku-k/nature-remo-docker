# syntax = docker/dockerfile:experimental
FROM golang:1.16-buster AS build
WORKDIR /usr/src
COPY . /usr/src/
RUN --mount=type=cache,target=/go \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -ldflags='-s -w'

FROM scratch
COPY --from=build /usr/src/natureremo-cli /remo
CMD ["/remo"]
