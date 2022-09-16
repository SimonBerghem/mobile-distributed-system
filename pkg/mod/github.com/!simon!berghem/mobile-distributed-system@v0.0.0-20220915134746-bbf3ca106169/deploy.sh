#!/bin/bash

# Run this file from project root folder with sudo sh ./deploy.sh
docker stack rm kadlab
docker build . -t kadlab
docker stack deploy nodes --compose-file docker-compose.yml
docker ps
