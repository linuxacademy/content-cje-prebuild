FROM golang:1.13

WORKDIR /go/src
COPY . .


CMD ["makeindex"]

