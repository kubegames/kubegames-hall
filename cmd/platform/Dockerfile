FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD platform /platform
RUN chmod -R 777 /platform

WORKDIR /

ENTRYPOINT ./platform