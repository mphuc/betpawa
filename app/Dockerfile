FROM golang:1.14.4-alpine as builder
RUN mkdir -p /a/app
COPY app.go /a/app
COPY go.mod /a/app
COPY go.sum /a/app
ENV GO111MODULE=on 
WORKDIR /a/app
RUN go install 

FROM alpine:3.10
# for testing 
RUN apk add curl
CMD ["./betpawa"]
COPY --from=builder /go/bin/betpawa .
