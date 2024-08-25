FROM golang:1.23.0-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go-demo

CMD [ "/go-demo" ]