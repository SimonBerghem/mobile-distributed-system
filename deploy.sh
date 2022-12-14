#!/bin/bash

# Run this file from project root folder with sudo sh ./deploy.sh
#docker stack rm nodes
docker build . -t kadlab
# docker stack deploy nodes --compose-file docker-compose.yml
COMPOSE_HTTP_TIMEOUT=200 docker-compose --compatibility up -d
docker ps
