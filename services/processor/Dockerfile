FROM golang:1.15.0

WORKDIR /go/src

COPY cmd ./cmd
COPY pkg ./pkg
COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build -o ./processor ./cmd/main.go

EXPOSE 10002

ENTRYPOINT ["./processor"]
