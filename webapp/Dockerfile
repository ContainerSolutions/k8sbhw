FROM golang:1.9 as BUILDER

WORKDIR /go/src/github.com/riccardomc/k8sbhw/webapp

# Build and test deps
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/postgres
RUN go get github.com/jinzhu/gorm/dialects/sqlite

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
