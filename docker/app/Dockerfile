ARG GO_VERSION=1.21.1
ARG ALPINE_VERSION=3.18

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /app

RUN apk update \
    && apk add --no-cache git \
    && apk add bash 

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download