#!/bin/bash

# Build Docker images for local development
echo "Building Docker images..."

# Build state-manager
echo "Building state-manager..."
docker build -t plarail2023/state-manager:local -f docker/backend/state-manager/Dockerfile .

# Build auto-operation
echo "Building auto-operation..."
docker build -t plarail2023/auto-operation:local -f docker/backend/auto-operation/Dockerfile .

# Build frontend
echo "Building frontend..."
docker build -t plarail2023/frontend:local -f docker/frontend/dashboard/Dockerfile.static .

# Build proxy
echo "Building proxy..."
docker build -t plarail2023/proxy:local -f docker/backend/proxy/Dockerfile .

echo "All images built successfully!"
