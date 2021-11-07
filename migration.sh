#!/bin/sh

export MIGRATION_DIR=migration/
export DB_PORT="5432"
export DB_HOST="localhost"
export DB_NAME="sqlc_test"
export DB_USER="postgres"
export DB_SSL=disable
export DB_PASSWORD="qwerty"
export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL} binary_parameters=yes"

if [ "$1" = "--dryrun" ]; then
    goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" status -v
elif [ "$1" = "--reset" ]; then
    goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" reset -v
else
    goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v
fi