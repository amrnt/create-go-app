FROM golang:1.8.1-alpine
RUN apk add --update git make && rm -rf /var/cache/apk/*

WORKDIR /go/src/github.com/amrnt/create-go-app
ADD . /go/src/github.com/amrnt/create-go-app

RUN make vendor
RUN make install

CMD create-go-app server
