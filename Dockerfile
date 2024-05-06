FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go mod tidy

RUN go mod download

RUN go build -o /app

EXPOSE 8080

CMD [ "/app"]