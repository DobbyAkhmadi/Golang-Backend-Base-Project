#!/bin/bash
# This script deploys the microservices using Kubernetes

SERVICE_NAME="service1" # Replace with your service name
K8S_NAMESPACE="your-namespace"

# Deploy the microservice to Kubernetes
kubectl apply -f ./deployments/kubernetes/"$SERVICE_NAME"/deployment.yaml -n "$K8S_NAMESPACE"

echo "$SERVICE_NAME deployment complete."
