FROM golang:1.13 as builder

RUN mkdir /go/src/github.com

ADD . /go/src/github.com/calculator

WORKDIR /go/src/github.com/calculator

RUN go test ./... && \
CGO_ENABLED=0 go build -o calc .

# --------------------------------------------------------------
FROM golang:1.13-alpine

ENV GIN_MODE=release

RUN mkdir /go/src/app

COPY --from=builder /go/src/github.com/calculator/calc /go/src/app

CMD ["/go/src/app/calc"]