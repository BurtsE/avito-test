FROM golang:1.22-alpine AS builder

WORKDIR /build



ADD go.mod .
ADD go.sum .
RUN go mod download
RUN pwd


COPY cmd cmd
COPY internal internal

RUN GOOS=linux go build -o app ./cmd

FROM golang:1.22-alpine
WORKDIR /root/

COPY configs configs

ENV HOUSE-DB_USER=${HOUSE-DB_USER}
ENV HOUSE-DB_PASSWORD=${HOUSE-DB_PASSWORD}
ENV USER-DB_USER=${HOUSE-DB_USER}
ENV USER-DB_PASSWORD=${HOUSE-DB_PASSWORD}
COPY --from=builder /build/app .


CMD ["./app"]