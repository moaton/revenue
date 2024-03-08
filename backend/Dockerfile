FROM golang:latest
RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update

RUN go mod download
RUN go build -o backend ./cmd/main.go

CMD ["./backend/"]