#/bin/sh

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$DB_NAME" -c "
    CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
    CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\" WITH SCHEMA public;"