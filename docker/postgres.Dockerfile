FROM postgres:latest
COPY ../db/init.sql /docker-entrypoint-initdb.d/