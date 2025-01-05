FROM golang:1.23.4-alpine 

WORKDIR /app
COPY go.mod .
RUN go mod tidy
COPY main.go .
