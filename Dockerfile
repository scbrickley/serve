FROM golang:latest AS go

WORKDIR /src

COPY go.mod .
COPY main.go .

EXPOSE 8081

RUN go build -o serve

ENTRYPOINT ["./serve"]
