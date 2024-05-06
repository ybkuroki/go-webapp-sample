FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build github.com/ybkuroki/go-webapp-sample/

EXPOSE 8080

CMD [ "/master" ]
