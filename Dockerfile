FROM golang:1.17.2 as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build .

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /app/Upload-Service .
RUN mkdir uploads

ENTRYPOINT ["./Upload-Service"]