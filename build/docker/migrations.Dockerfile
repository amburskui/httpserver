ARG MIGRATE_VERSION=v4.15.0
FROM migrate/migrate:${MIGRATE_VERSION}
WORKDIR /app
COPY [ "migrations/", "./" ]