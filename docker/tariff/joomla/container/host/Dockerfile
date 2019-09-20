FROM ubuntu:18.04

ARG SSH_ROOT_PASS

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    vim \
    sudo \
    openssh-server

COPY user.sh /usr/local/bin/user.sh
RUN chmod +x /usr/local/bin/user.sh
RUN sed -i 's/{secret}/$SSH_ROOT_PASS/' /usr/local/bin/user.sh
RUN /usr/local/bin/user.sh

COPY sshd_config /etc/ssh/sshd_config

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]

CMD tail -f /dev/null