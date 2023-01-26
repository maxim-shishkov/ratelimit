FROM golang:latest

WORKDIR /app
COPY ./ /app

RUN go mod download

RUN cd /app && GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .

EXPOSE 8000

ENTRYPOINT go run .