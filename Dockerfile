FROM golang:1.8-alpine AS builder

MAINTAINER Jose Maria Hidalgo Garcia <jhidalgo3@gmail.com>

RUN apk add --no-cache --update ca-certificates \
    && apk add curl git coreutils \
    && rm /var/cache/apk/*

ENV SRC_DIR=/go/src/github.com/jhidalgo3/training-docker-microservice

CMD go get github.com/Masterminds/glide

ADD ./src/github.com/jhidalgo3/training-docker-microservice ${SRC_DIR}

WORKDIR ${SRC_DIR}

RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/jhidalgo3/training-docker-microservice/training-docker-microservice  .

CMD ./training-docker-microservice