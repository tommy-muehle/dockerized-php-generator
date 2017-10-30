FROM alpine:3.6
MAINTAINER Tommy Muehle <tommy.muehle@gmail.com>

RUN apk --update --progress --no-cache --repository http://dl-cdn.alpinelinux.org/alpine/edge/community add \
    php7 \
    php7-xdebug \
    php7-openssl \
    php7-ctype \
    php7-json \
    php7-phar \
    php7-zlib \
    php7-mbstring \
    php7-tokenizer \
    php7-dom \
    php7-xml \
    php7-xmlwriter

VOLUME /app
WORKDIR /app

ENTRYPOINT ["tool"]
