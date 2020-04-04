FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 7777

ENTRYPOINT ["./main", "config.json"]

