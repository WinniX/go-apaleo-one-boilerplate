#/bin/sh

psql <<- EOSQL
    CREATE DATABASE $DB_NAME;
    CREATE USER $DB_USER WITH LOGIN PASSWORD '$DB_PASSWORD' CREATEDB;
    GRANT ALL PRIVILEGES ON DATABASE ${DB_NAME} to $POSTGRES_USER;
EOSQL
