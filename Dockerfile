FROM golang:1.20.2

COPY ./app /go-gin-init-v2/app
COPY ./go.mod /go-gin-init-v2/

WORKDIR /go-gin-init-v2

RUN GOOS=linux GOARCH=amd64 go build ./app/cmd/main.go ./app/cmd/wire_gen.go

CMD ["./main"]
