FROM golang:1.19-alpine as builder
ARG BUILD_DIR=/webapp
# libc6-compat is required for building go binaries in Alpine
RUN apk --update add --no-cache gcc libc-dev && mkdir $BUILD_DIR
WORKDIR $BUILD_DIR
COPY . $BUILD_DIR
RUN GOOS=linux go build -o $BUILD_DIR main.go

FROM alpine:3.18 as production
COPY --from=builder /webapp .
CMD ["./main"]