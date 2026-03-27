#!/usr/bin/env bash

set -euo pipefail
umask 077

export PG_PASSWORD_DAVID="$(openssl rand -base64 15)"
export PG_PASSWORD_GOKANBAN="$(openssl rand -base64 15)"

echo "export PG_PASSWORD_DAVID=\"$PG_PASSWORD_DAVID\""
echo "export PG_PASSWORD_GOKANBAN=\"$PG_PASSWORD_GOKANBAN\""

envsubst \
  < migrations/000001_create_roles.up.sql \
  > migrations/000001_create_roles.up.sql.temp

mv \
  migrations/000001_create_roles.up.sql.temp \
  migrations/000001_create_roles.up.sql
