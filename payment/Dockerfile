FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o payment .

EXPOSE 8084

CMD ["./payment"]
