FROM golang:1.21-alpine3.19 as build

WORKDIR /go/src/webapp

COPY main.go main.go

RUN CGO_ENABLED=0 GOOS=linux go build -o webapp main.go

FROM alpine:3.19

COPY --from=build /go/src/webapp /app
COPY ./static /app/static

WORKDIR /app

RUN chmod 777 ./static/index.html

EXPOSE 3000

ENTRYPOINT ./webapp
