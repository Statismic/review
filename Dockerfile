FROM golang:alpine

WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build

EXPOSE 8000
CMD ["./app"]