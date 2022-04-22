FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD game /game
RUN chmod -R 777 /game

WORKDIR /

ENTRYPOINT ./game