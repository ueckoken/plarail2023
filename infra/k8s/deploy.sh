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

# Apply the manifests using kustomize
echo "Applying Kubernetes manifests..."
kubectl apply -k overlays/local/

echo ""
echo "Deployment complete!"
echo ""
echo "Services are available at:"
echo "- Frontend: http://localhost:30031 (via proxy) or use Ingress"
echo "- EMQX Dashboard: http://localhost:31808"
echo "- Mongo Express: http://localhost:30081"
echo "- MQTT: localhost:31883"
echo ""
echo "If using Ingress with NGINX Ingress Controller:"
echo "- Frontend: http://localhost/"
echo "- API: http://localhost/api/"
echo "- EMQX Dashboard: http://localhost/emqx/"
echo "- Mongo Express: http://localhost/mongo-express/"
