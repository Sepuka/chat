#!/bin/bash
set -e

printf "\n\033[0;44m---> check root passwd.\033[0m\n"

export randompasswd=$(date +%s | sha256sum | base64 | head -c 32 ; echo)
printf "\n\033[0;44m--->  root passwd $randompasswd .\033[0m\n"
echo "root:secret" | chpasswd

printf "\n\033[0;44m---> Starting the nginx server.\033[0m\n"

service nginx start

printf "\n\033[0;44m---> Starting the SSH server.\033[0m\n"

service ssh start
service ssh status

exec "$@"