FROM centos:7
MAINTAINER <andy@cohan.org>

RUN yum update -y && \
  yum install -y nc

WORKDIR /app
COPY *.sh /app/

CMD ["/app/entrypoint.sh"]
