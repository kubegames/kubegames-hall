FROM busybox:latest

ADD Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

ADD single /single
RUN chmod -R 777 /single

WORKDIR /

ENTRYPOINT ./single