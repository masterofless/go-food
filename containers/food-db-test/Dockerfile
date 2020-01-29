FROM golang:latest
MAINTAINER <andy@cohan.org>

WORKDIR /app
COPY *.sh /app/
COPY cmd/fooddb_test/* /app/

RUN go mod init github.com/masterofless/go-food/food-db-test
RUN go get github.com/couchbase/gocb/v2@v2.0.0-beta.1
RUN go test -v .

CMD ["/app/entrypoint.sh"]
