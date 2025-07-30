FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /server .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=build /server .

COPY .env .

EXPOSE 8080

CMD ["./server"]
