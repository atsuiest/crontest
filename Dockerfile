FROM golang:1.13-alpine AS builder 

RUN apk update && apk add git

RUN adduser -D -g ‘’ appuser

# COPY . $GOPATH/src/bupa.cl

WORKDIR $GOPATH/src/crontest


COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/crontest

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

USER appuser

COPY --from=builder /go/bin/crontest /go/bin/crontest


#COPY . /go/bin/

CMD ["/go/bin/crontest"]