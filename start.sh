#!bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -databasee "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"


