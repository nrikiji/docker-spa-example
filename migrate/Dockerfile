FROM golang:1.17.7-alpine3.15 AS build

ARG CGO_ENABLED=0

WORKDIR /go/src/app
RUN go get github.com/rubenv/sql-migrate/...

COPY . .
CMD ["sql-migrate", "up", "-env", "production"]