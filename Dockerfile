FROM golang:1.21.6-alpine AS builder

WORKDIR /app

COPY ./ ./

RUN mkdir bin
RUN go mod tidy
RUN go build -o ./bin/carApp ./server.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/.env ./
COPY --from=builder /app/bin/carApp ./

EXPOSE 5005

CMD ./carApp