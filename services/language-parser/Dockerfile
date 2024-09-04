FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /language-parser

EXPOSE 3000

CMD ["go", "run", "main.go"]
