FROM centos:7

RUN yum update -y && yum install -y \
    jq \
    curl

WORKDIR /app
COPY entrypoint.sh /app/

CMD [ "/app/entrypoint.sh" ]
