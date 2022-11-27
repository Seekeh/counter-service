FROM golang:1.17

WORKDIR /app

COPY main.go ./

EXPOSE 8080

CMD ["go", "run", "main.go"]
