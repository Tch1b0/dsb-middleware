FROM golang:latest

WORKDIR /app

COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN go mod download
RUN go build -o main

CMD ["./main"]