FROM golang:1.10

EXPOSE 8080

WORKDIR /go/src/gw
COPY src/gw .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["gw"]
