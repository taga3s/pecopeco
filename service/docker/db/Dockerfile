ARG MYSQL_VERSION=8.0
FROM --platform=linux/amd64 mysql:${MYSQL_VERSION}-debian

RUN apt-get update \
    && apt-get -y install locales-all

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8

ENV TZ Asia/Tokyo

COPY ./service/docker/db/conf/my.cnf /etc/my.cnf