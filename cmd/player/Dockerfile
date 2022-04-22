FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD player /player
RUN chmod -R 777 /player

WORKDIR /

ENTRYPOINT ./player