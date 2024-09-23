FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o order .

EXPOSE 8083

CMD ["./order"]
