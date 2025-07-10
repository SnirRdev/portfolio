FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["go run main.go"]
