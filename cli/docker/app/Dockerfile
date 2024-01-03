ARG GO_VERSION=1.21.1
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /app

RUN apk update \
    && apk add --no-cache git \
    && apk add bash 

RUN go install github.com/spf13/cobra-cli@latest

COPY ./cli/go.mod .
COPY ./cli/go.sum .

RUN go mod download