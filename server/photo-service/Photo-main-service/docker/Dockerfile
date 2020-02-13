FROM golang:latest

COPY . /usr/src/myapp
WORKDIR /usr/src/myapp

ENV GOOGLE_APPLICATION_CREDENTIALS="/usr/src/myapp/Final-Year-GMIT-329249b5df19.json"


EXPOSE 30051

ENTRYPOINT ["./main"]

