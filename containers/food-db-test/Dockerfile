FROM golang:1.13
MAINTAINER <andy@cohan.org>

WORKDIR /app

#RUN go mod init github.com/masterofless/go-food/food-db-test
RUN go mod init go-food/food-db-test
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo

COPY *.sh /app/
COPY cmd/food-db-test/* /app/
RUN go build

CMD ["/app/entrypoint.sh"]
