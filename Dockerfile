
FROM golang:1.15.1-alpine3.12

WORKDIR /go/sakti-app-role

COPY . /go/sakti-app-role

RUN go build main.go

EXPOSE 10000

CMD ["./main"]