FROM golang:1.24.0-alpine AS builder

WORKDIR /usr/local/src

RUN apk add --no-cache build-base

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

RUN go build -o bin/app

# Этап запуска
FROM alpine AS runner

WORKDIR /app

RUN apk add --no-cache sqlite ca-certificates

COPY --from=builder /usr/local/src/bin/app ./app

CMD ["./app"]
