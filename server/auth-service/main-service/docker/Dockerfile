FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 50051

ENTRYPOINT ["./main", "config.json"]

