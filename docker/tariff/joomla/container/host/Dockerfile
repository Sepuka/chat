FROM ubuntu:18.04

ARG SSH_ROOT_PASS
ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    vim \
    sudo \
    unzip \
    php7.2-fpm \
    php7.2-xml \
    php7.2-mysql \
    openssh-server

RUN apt-get install -y mysql-server
RUN apt-get install -y nginx

ADD mysqld.conf /etc/mysql/my.cnf
ADD init.sh /usr/local/bin/init.sh
RUN chmod +x /usr/local/bin/init.sh
RUN chown -R mysql:mysql /var/lib/mysql

COPY joomla.nginx.conf /etc/nginx/sites-available/default

COPY sshd_config /etc/ssh/sshd_config

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

COPY Joomla_3.9.14-Stable-Full_Package.zip /var/www/joomla.zip
RUN unzip -qq /var/www/joomla.zip -d /var/www/joomla
RUN rm -rf /var/www/joomla.zip

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD tail -f /dev/null