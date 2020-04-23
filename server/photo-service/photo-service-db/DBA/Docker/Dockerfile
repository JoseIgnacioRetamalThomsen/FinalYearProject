FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 7172

ENTRYPOINT ["./pdba", "config.json"]

