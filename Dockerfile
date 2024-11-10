FROM ubuntu:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn

WORKDIR /app

COPY . .

RUN go build -o main

ENV TZ=Asia/Shanghai\
    LANG=zh_CN.utf8

EXPOSE 8080

CMD ["./main"]