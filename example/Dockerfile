FROM golang:1.19.5-alpine3.17 AS base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o provider ./cmd/provider
RUN go build -o consumer ./cmd/consumer

FROM alpine:3.17 AS runner

WORKDIR /app

COPY --from=base /build/provider ./provider
COPY --from=base /build/consumer ./consumer