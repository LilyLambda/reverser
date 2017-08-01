FROM golang:1.8.1-alpine

RUN apk --no-cache add git

ENV GOROOT /usr/local/go/
ENV GOPATH /go

WORKDIR /go/src/github.com/lilylambda/reverser/
COPY *.go /go/src/github.com/lilylambda/reverser/
COPY cmd /go/src/github.com/lilylambda/reverser/cmd/
COPY vendor /go/src/github.com/lilylambda/reverser/vendor/
RUN go install \
    github.com/lilylambda/reverser \
    github.com/lilylambda/reverser/cmd/reverser

ENTRYPOINT ["/go/bin/reverser"]
