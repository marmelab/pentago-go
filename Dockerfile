FROM golang:alpine

ENV CGO_ENABLED 0
WORKDIR $GOPATH
