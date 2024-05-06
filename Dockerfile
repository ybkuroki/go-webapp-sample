FROM golang:1.22-alpine

WORKDIR .

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /master

EXPOSE 8080

CMD [ "/master" ]
