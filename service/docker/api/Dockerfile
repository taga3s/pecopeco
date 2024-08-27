ARG GO_VERSION=1.23.0
ARG ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}

WORKDIR /app

RUN apk update \
    && apk add --no-cache git \
    && apk add bash 

COPY ./service/go.mod .
COPY ./service/go.sum .

RUN go mod download
