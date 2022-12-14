#!/bin/bash

# Wait until Postgres is ready.
while ! pg_isready -q -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER
do
  echo "$(date) - waiting for database to start - $POSTGRES_HOST $POSTGRES_PORT $POSTGRES_USER $POSTGRES_DB"
  sleep 2
done

# Create, migrate, and seed database if it doesn't exist.
if [[ -z `psql -Atqc "\\list $POSTGRES_DB"` ]]; then
  echo "Database $POSTGRES_DB does not exist. Creating..."
  createdb -E UTF8 $POSTGRES_DB -l en_US.UTF-8 -T template0
  echo "Database $POSTGRES_DB created."
fi


