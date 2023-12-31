#!/bin/bash
#set -e

POSTGRES="psql --username ${APPLICATION_POSTGRES_USER}"

echo "Creating database role: ${DB_USER}"

$POSTGRES <<-EOSQL
CREATE USER ${DB_USER} WITH CREATEDB PASSWORD '${DB_PASS}';
EOSQL
