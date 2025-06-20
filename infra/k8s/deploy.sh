#!/bin/bash

# Deploy to Kubernetes using Kustomize
echo "Deploying to Kubernetes..."

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "kubectl could not be found. Please install kubectl first."
    exit 1
fi

# Check if the cluster is accessible
if ! kubectl cluster-info &> /dev/null; then
    echo "Cannot connect to Kubernetes cluster. Please check your kubeconfig."
    exit 1
fi

# Check if nginx-ingress-controller is installed
echo "Checking nginx-ingress-controller..."
if ! kubectl get deployment -n plarail nginx-ingress-ingress-nginx-controller &> /dev/null; then
    echo "nginx-ingress-controller is not installed."
    echo "Installing nginx-ingress-controller..."
    ./install-nginx-ingress.sh
    if [ $? -ne 0 ]; then
        echo "Failed to install nginx-ingress-controller"
        exit 1
    fi
else
    echo "nginx-ingress-controller is already installed."
fi

# Apply the manifests using kustomize
echo ""
echo "Applying Kubernetes manifests..."
kubectl apply -k overlays/local/

echo ""
echo "Deployment complete!"
echo ""
echo "Services are available at:"
echo "- Frontend: http://localhost:30080 (via Ingress)"
echo "- EMQX Dashboard: http://localhost:31808"
echo "- Mongo Express: http://localhost:30081"
echo "- MQTT: localhost:31883"
echo "- Proxy: http://localhost:30031"
echo ""
echo "Ingress endpoints (via http://localhost:30080):"
echo "- Frontend: http://localhost:30080/"
echo "- API: http://localhost:30080/api/"
echo "- EMQX Dashboard: http://localhost:30080/emqx/"
echo "- Mongo Express: http://localhost:30080/mongo-express/"
echo "- Proxy: http://localhost:30080/proxy/"
