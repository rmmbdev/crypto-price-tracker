FROM golang:1.22-alpine as builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o bin/migrator ./migrator
RUN go build -o bin/updater ./updater

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /build/bin/migrator .
COPY --from=builder /build/bin/updater .