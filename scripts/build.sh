#!/bin/bash
# This script builds the microservices

SERVICE_NAME="service1" # Replace with your service name
OUTPUT_DIR="./build"

echo "Building $SERVICE_NAME..."
go build -o "$OUTPUT_DIR/$SERVICE_NAME" ./cmd/"$SERVICE_NAME"

echo "$SERVICE_NAME build complete."
