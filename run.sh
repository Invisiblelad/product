#!/bin/bash

docker stop product_container || true
docker rm product_container || true

echo "Building the Docker image..."
docker build -t product_image .

echo "Running the Docker container..."
docker run -d -p 8080:8080 -e MONGO_URI=mongodb://172.23.212.112:27017 --name product_container product_image

docker ps | grep product_container

