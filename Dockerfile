FROM golang:alpine AS builder

WORKDIR /app


COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd/mokhtasar
RUN go build -o /cmd

FROM alpine:latest


WORKDIR /app/

COPY --from=builder /cmd .

EXPOSE 5000

ENTRYPOINT ["./cmd"]

CMD ["server"]
