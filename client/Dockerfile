FROM golang:1.13-alpine

WORKDIR /usr/src
COPY . .
RUN GOOS=linux go build

FROM alpine:latest
WORKDIR /root
COPY --from=0 /usr/src/natureremo-cli /usr/bin/remo
