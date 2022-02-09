FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on 
RUN go get -u github.com/ybkuroki/go-webapp-sample
RUN cd /build && git clone https://github.com/AdminTurnedDevOps/go-webapp-sample.git

# Download all the dependencies
RUN cd /build/go-webapp-sample && go get -d -v ./...

# Install the package
RUN cd /build/go-webapp-sample && go install -v ./...

RUN cd /build/go-webapp-sample && go build

EXPOSE 8080

ENTRYPOINT [ "/build/go-webapp-sample/go-webapp-sample" ]