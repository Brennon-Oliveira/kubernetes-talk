FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go build -o product .

EXPOSE 8082

CMD ["./product"]
