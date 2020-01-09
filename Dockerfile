FROM golang:1.13 as builder

ENV DEP_VERSION=0.5.4

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 && \
chmod +x /usr/local/bin/dep && \
mkdir /go/src/github.com

ADD . /go/src/github.com/calculator

WORKDIR /go/src/github.com/calculator

RUN dep ensure -vendor-only && \
go test ./... && \
CGO_ENABLED=0 go build -o calc .

# --------------------------------------------------------------
FROM golang:1.13-alpine

ENV GIN_MODE=release

RUN mkdir /go/src/app

COPY --from=builder /go/src/github.com/calculator/calc /go/src/app

CMD ["/go/src/app/calc"]