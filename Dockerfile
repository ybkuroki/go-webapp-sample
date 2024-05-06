FROM golang:1.22-alpine

WORKDIR alhusseinahmed/go-webapp-sample

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build /config
RUN go build /container
RUN go build /logger
RUN go build /middleware
RUN go build /migration
RUN go build /repository
RUN go build /router
RUN go build /session

EXPOSE 8080

CMD [ "/master" ]
