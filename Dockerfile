FROM golang:1.10.4-alpine3.8

RUN apk update
RUN apk add git
RUN apk add alpine-sdk

WORKDIR /go/src/bitbucket.org/242617/blockchain.com-bruteforce
COPY . .

RUN go get -d -v ./...
RUN make

CMD ["sh", "-c", "cp build/* ./"]