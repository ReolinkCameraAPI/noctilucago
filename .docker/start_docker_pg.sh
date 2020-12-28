#!/bin/bash
# sudo rm -rf postgres/pgdata
# mkdir postgres/pgdata
#UserID=${UID} GroupID=${GID} \
docker-compose -f ./postgres/docker-compose.pg.yml up \
--build \
--force-recreate \
--remove-orphans
