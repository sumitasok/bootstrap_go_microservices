FROM golang:1.9

WORKDIR /go/src/app
COPY . .

# requires cgo to work
# export CGO_ENABLED=1

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]