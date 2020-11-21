FROM alpine:latest

# 更新安装源
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache bash supervisor \
    && mkdir /app
WORKDIR /app
ADD demo-service /app/demo-service
ADD supervisord.conf /etc/supervisord.conf

#CMD ["/app/demo-service"]
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]

