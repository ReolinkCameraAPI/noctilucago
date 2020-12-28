# Dockerfile
FROM postgres:12.4-alpine
MAINTAINER Alano Terblanche "alanoterblanche@gmail.com"
USER postgres
COPY init.sh /docker-entrypoint-initdb.d/