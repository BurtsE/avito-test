FROM golang:1.21-alpine AS builder

WORKDIR /build


ADD go.mod .
ADD go.sum .
RUN go mod download
RUN pwd


COPY cmd cmd
COPY internal internal

RUN GOOS=linux go build -o app ./cmd

FROM golang:1.21-alpine
WORKDIR /root/

COPY configs configs

ENV PG_USER=${HOUSE-DB_USER}
ENV PG_PASSWORD=${HOUSE-DB_PASSWORD}
COPY --from=builder /build/app .


CMD ["./app"]