FROM golang:latest

RUN mkdir -p /go/src/kitsu-rest
WORKDIR /go/src/kitsu-rest

COPY . /go/src/kitsu-rest

RUN go get -d -v
RUN go install -v

EXPOSE 1812

CMD ["/go/bin/kitsurest"]
