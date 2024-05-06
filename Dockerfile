FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go get github.com/ybkuroki/go-webapp-sample/config
RUN go build github.com/ybkuroki/go-webapp-sample/config
RUN go build github.com/ybkuroki/go-webapp-sample/container
RUN go build github.com/ybkuroki/go-webapp-sample/logger
RUN go build github.com/ybkuroki/go-webapp-sample/middleware
RUN go build github.com/ybkuroki/go-webapp-sample/migration
RUN go build github.com/ybkuroki/go-webapp-sample/repository
RUN go build github.com/ybkuroki/go-webapp-sample/router
RUN go build github.com/ybkuroki/go-webapp-sample/session

EXPOSE 8080

CMD [ "/master" ]
