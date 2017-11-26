FROM golang:1.9 as BUILDER

WORKDIR /go/src/github.com/riccardomc/k8sbhw/webapp

COPY webapp.go .
COPY webapp_test.go .
COPY datastore ./datastore

RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o webapp -a -ldflags '-extldflags "-static"' .

# Run

FROM scratch
COPY --from=BUILDER /go/src/github.com/riccardomc/k8sbhw/webapp/webapp /webapp
EXPOSE 9009
ENTRYPOINT ["/webapp"]
