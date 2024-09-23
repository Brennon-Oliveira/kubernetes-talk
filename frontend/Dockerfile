FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o frontend .

EXPOSE 8080

CMD ["./frontend"]