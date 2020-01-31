FROM golang:1.13
MAINTAINER <andy@cohan.org>

WORKDIR /app
COPY *.sh /app/
COPY cmd/food-api/* /app/

#RUN go mod init github.com/masterofless/go-food/food-api
RUN go mod init go-food/food-api
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go test -v .

CMD ["/app/entrypoint.sh"]
