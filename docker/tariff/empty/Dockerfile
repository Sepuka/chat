FROM ubuntu:18.04

ARG SSH_ROOT_PASS
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    vim \
    sudo \
    php7.2-fpm \
    openssh-server

RUN apt-get install -y nginx

COPY sshd_config /etc/ssh/sshd_config

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD tail -f /dev/null