#build stage
FROM golang:alpine AS builder
# change apk source to ali
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update && apk upgrade
# install: git build-base opencv-dev
RUN apk add --no-cache git build-base
# change go proxy to qiniu cdn
RUN go env -w GOPROXY=https://goproxy.cn,direct
# build app
WORKDIR /go/src/app
COPY . .
# install gocv
# RUN go get -u -d gocv.io/x/gocv && go install gocv.io/x/gocv
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./client/main.go

#final stage
FROM alpine:latest
# change apk source to ali
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && apk update && apk upgrade
# install: tzdata ca-certificates opencv-dev
RUN apk --no-cache add tzdata ca-certificates
# cpoy the app from builder
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=cocktail98/jetson-nano-detect-person-service-go Version=1.0
EXPOSE 50001 50002 6380 6381
