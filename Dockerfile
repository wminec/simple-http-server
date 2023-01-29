FROM golang:alpine AS build

RUN apk add git

RUN mkdir /src
ADD . /src
WORKDIR /src

run go build -o /tmp/simple-http-server ./cmd/simple-http-server/main.go

FROM alpine:edge
COPY --from=build /tmp/simple-http-server /sbin/http-server

CMD /sbin/http-server