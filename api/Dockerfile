FROM golang:1.24-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o . ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]