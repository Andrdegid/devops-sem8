#!/bin/bash

echo "Сборка образа из Dockerfile.system..."
docker build -f Dockerfile.system -t system .

echo "Сборка образа из Dockerfile.builder..."
docker build -f Dockerfile.builder -t build .

echo "Запуск docker-compose..."
docker-compose up -d