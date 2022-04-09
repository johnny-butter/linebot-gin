#!/bin/sh

set -ex

MIGRATIONS_PATH="models/migrations"

case $1 in
  "up" )
    migrate -database ${MIGRATIONS_DB_URL} -path ${MIGRATIONS_PATH} up $2
;;
  "down" )
    migrate -database ${MIGRATIONS_DB_URL} -path ${MIGRATIONS_PATH} down $2
;;
  "new" )
    migrate create -ext sql -dir ${MIGRATIONS_PATH} -seq $2
;;
  "force" )
    migrate -database ${MIGRATIONS_DB_URL} -path ${MIGRATIONS_PATH} force $2
;;
esac
