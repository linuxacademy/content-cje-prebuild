FROM golang:1.13

WORKDIR /go
COPY ./src .

RUN go build ./makeindex.go
RUN ./makeindex 

