FROM golang:1.22.1-alpine3.19 as build

COPY . /app

WORKDIR /app

RUN apk add make && make build

#====================================

FROM alpine:3.19

WORKDIR /

# COPY /deploy/migrations /migrations
COPY --from=build /app/bin/ /bin

RUN mkdir /logs


CMD [ "bin/app" ]
