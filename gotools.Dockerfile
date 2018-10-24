FROM golang:1.9

WORKDIR /go/src/app

# add your tools here.
RUN go get -v github.com/rubenv/sql-migrate/...
RUN go get -u github.com/golang/dep/cmd/dep

CMD [ "bash" ]