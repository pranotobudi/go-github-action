FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY static ./static

RUN go build -o /go-github-action

EXPOSE 8000

CMD ["/go-github-action"]
