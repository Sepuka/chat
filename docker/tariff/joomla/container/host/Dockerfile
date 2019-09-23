FROM ubuntu:18.04

ARG SSH_ROOT_PASS

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    vim \
    sudo \
    openssh-server

RUN apt-get install -y mysql-server
RUN apt-get install -y nginx
ADD mysqld.conf /etc/mysql/my.cnf

COPY joomla.nginx.conf /etc/nginx/sites-available/default

COPY sshd_config /etc/ssh/sshd_config

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD tail -f /dev/null