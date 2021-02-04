FROM alpine:3.6

RUN apk add --no-cache \
        ca-certificates \
        bash \
    && rm -f /var/cache/apk/*

COPY bin/productmanagement /usr/local/bin/todoapp

CMD ["/usr/local/bin/todoapp"]
