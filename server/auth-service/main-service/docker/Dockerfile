FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 50051

ENTRYPOINT ["./c/main", "config.json"]

