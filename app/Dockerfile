FROM golang:1.22-alpine

RUN addgroup -S appuser && adduser -S appuser -G appuser

WORKDIR /app

COPY --chown=appuser:appuser . /app

USER appuser
# go will be replaced later
RUN go build -o main main.go  

EXPOSE 8080

ENTRYPOINT ["./main"]

