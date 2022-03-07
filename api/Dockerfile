FROM golang:1.17.7-alpine3.15 AS build

WORKDIR /go/src/app

ADD ./go.mod .
ADD ./go.sum .
RUN go mod download

COPY . .
RUN go build -o server .

FROM alpine:3.15.0
WORKDIR /go/src/app

COPY .env .
COPY --from=build /go/src/app/server .

EXPOSE 1323
CMD ["./server"]
