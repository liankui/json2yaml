FROM golang:1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o json2yaml

FROM alpine:3.18.9
ENV LANG C.UTF-8
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk update && apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && apk del tzdata
WORKDIR /root/
COPY --from=builder /app/json2yaml .
COPY --from=builder /app/static ./static
CMD ["./json2yaml"]
EXPOSE 21111
