#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER dockeruser;
    GRANT ALL PRIVILEGES ON DATABASE goboilerdb TO dockeruser;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO dockeruser;
EOSQL%