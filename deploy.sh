#!/bin/bash

# Run this file from project root folder

docker stack deploy nodes --compose-file docker-compose.yml
docker ps