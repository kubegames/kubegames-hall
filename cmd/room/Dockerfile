FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD room /room
RUN chmod -R 777 /room

WORKDIR /

ENTRYPOINT ./room