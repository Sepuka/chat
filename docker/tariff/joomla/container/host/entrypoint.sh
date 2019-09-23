#!/bin/bash
set -e


printf "\n\033[0;44m---> check root passwd.\033[0m\n"

export randompasswd=$(date +%s | sha256sum | base64 | head -c 32 ; echo)
printf "\n\033[0;44m--->  root passwd $randompasswd .\033[0m\n"
sed "s/root:*:18151:0:99999:7:::/root:$randompasswd:18151:0:99999:7:::/g" /etc/shadow

printf "\n\033[0;44m---> Starting the mysql server.\033[0m\n"

service mysql start

printf "\n\033[0;44m---> Starting the nginx server.\033[0m\n"

service nginx start

printf "\n\033[0;44m---> Starting the SSH server.\033[0m\n"

service ssh start
service ssh status

exec "$@"