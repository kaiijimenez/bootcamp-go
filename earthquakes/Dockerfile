FROM golang:1.11.2
ADD . /go/src/github.com/kaiijimenez/bootcamp-go/earthquakes
WORKDIR /go/src/github.com/kaiijimenez/bootcamp-go/earthquakes
EXPOSE 8000
RUN go get -u github.com/gorilla/mux && go get -u github.com/go-sql-driver/mysql

CMD["go", "run", "main.go"]