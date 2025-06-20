#!/bin/bash

# nginx-ingress-controllerをHelmでインストールするスクリプト

set -e

echo "=== nginx-ingress-controller インストールスクリプト ==="

# Helmがインストールされているか確認
if ! command -v helm &> /dev/null; then
    echo "Error: Helmがインストールされていません。"
    echo "Helmをインストールしてください: https://helm.sh/docs/intro/install/"
    exit 1
fi

# nginx-ingress Helm repositoryを追加
echo "1. nginx-ingress Helm repositoryを追加..."
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update

# plarail namespaceが存在するか確認
echo "2. plarail namespaceを確認..."
kubectl get namespace plarail &> /dev/null || kubectl create namespace plarail

# nginx-ingress-controllerをインストール
echo "3. nginx-ingress-controllerをインストール..."
helm upgrade --install nginx-ingress ingress-nginx/ingress-nginx \
  --namespace plarail \
  --set controller.service.type=NodePort \
  --set controller.service.nodePorts.http=30080 \
  --set controller.service.nodePorts.https=30443 \
  --set controller.ingressClassResource.default=true \
  --set controller.ingressClass=nginx \
  --wait

echo ""
echo "=== インストール完了 ==="
echo ""

# インストール状態を確認
echo "nginx-ingress-controllerの状態:"
kubectl get pods -n plarail -l app.kubernetes.io/name=ingress-nginx
echo ""
kubectl get svc -n plarail -l app.kubernetes.io/name=ingress-nginx

echo ""
echo "=== アクセス方法 ==="
echo "http://localhost:30080 でIngressにアクセスできます"
echo ""
echo "Ingressの状態を確認:"
kubectl get ingress -n plarail
