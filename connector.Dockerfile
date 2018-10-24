FROM golang:1.9

WORKDIR /go/src/app

RUN go get -u github.com/golang/dep/...

#Add dependency list
ADD ./Gopkg.lock /go/src/app
ADD ./Gopkg.toml /go/src/app

# Restore go dependencies
RUN dep ensure -vendor-only -v

COPY . .

# requires cgo to work
# export CGO_ENABLED=1

RUN go install -v ./...

# run from the installed binary
CMD ["app"]