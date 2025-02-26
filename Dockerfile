#creating multi-layer build image

#creating base image
FROM golang:1.22 as base
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o  main .

#second layer on top of base

#choosing distroless image for reducing the size and security
FROM gcr.io/distroless/static-debian12 
WORKDIR /root/
COPY --from=base /app/main .
EXPOSE 8080
CMD ["./main"]
