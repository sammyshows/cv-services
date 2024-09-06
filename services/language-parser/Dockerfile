FROM golang:1.23

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

# Install air for hot-reloading (ARM64 version for M2)
RUN go install github.com/air-verse/air@latest


WORKDIR /app

# Copy the Go module files first
COPY go.mod ./
COPY go.sum ./

# Download Go module dependencies
RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o /language-parser

EXPOSE 3000

CMD ["air"]
