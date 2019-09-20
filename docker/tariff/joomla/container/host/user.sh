#!/bin/bash
set -e

printf "\n\033[0;44m---> Creating SSH root user.\033[0m\n"

echo "root:{secret}" | chpasswd

exec "$@"