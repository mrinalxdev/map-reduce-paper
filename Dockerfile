FROM golang:1.21-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o master cmd/master/main.go
RUN go build -o worker cmd/worker/main.go
