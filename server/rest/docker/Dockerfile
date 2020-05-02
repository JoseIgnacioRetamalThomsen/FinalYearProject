FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

EXPOSE 9371

ENTRYPOINT ["./main", "config.json"]

