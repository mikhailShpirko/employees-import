FROM postgres:17.4-alpine3.21
COPY src/persistence/migrations/1.Init.sql /docker-entrypoint-initdb.d/