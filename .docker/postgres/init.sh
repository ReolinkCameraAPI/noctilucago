#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE DATABASE "onecloudassets";
  GRANT ALL PRIVILEGES ON DATABASE "onecloudassets" TO postgres;

  CREATE DATABASE "onecloudassets-kratos";
  GRANT ALL PRIVILEGES ON DATABASE "onecloudassets-kratos" TO postgres;
EOSQL