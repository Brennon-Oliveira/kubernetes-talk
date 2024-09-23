FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o notification .

EXPOSE 8085

CMD ["./notification"]
