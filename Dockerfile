FROM golang:1.19.6

WORKDIR /go/src/app

COPY . .

RUN go mod download
RUN go build -o /hex-arch-microservice-port-scanner main.go

EXPOSE 8080

CMD ["./hex-arch-microservice-port-scanner"]