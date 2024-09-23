FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o auth .

EXPOSE 8081

CMD ["./auth"]
