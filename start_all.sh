#!/bin/bash

set -e

SERVICES=("auth-service" "product-service" "checkout-service")

# Stop and remove old containers if running
for SERVICE in "${SERVICES[@]}"; do
  if docker ps -a --format '{{.Names}}' | grep -q "^${SERVICE}$"; then
    echo "Stopping existing container: $SERVICE"
    docker rm -f $SERVICE >/dev/null 2>&1
  fi
done

# Build images
for SERVICE in "${SERVICES[@]}"; do
  echo "Building $SERVICE..."
  docker build --build-arg SERVICE=$SERVICE -t $SERVICE:local .
done

# Start containers on separate ports
echo "Starting containers..."

docker run -d --name auth-service -p 8080:8080 auth-service:local
docker run -d --name product-service -p 8081:8080 product-service:local
docker run -d --name checkout-service -p 8082:8080 checkout-service:local

echo ""
echo "Services running:"
echo "Auth:     http://localhost:8080"
echo "Products: http://localhost:8081"
echo "Checkout: http://localhost:8082"