#!/bin/bash

# Seed Data実行スクリプト
# 使用方法: ./seed-data.sh [環境] [データファイル]
# 例: ./seed-data.sh local backend/onetime/seed-data/data/mfk-2024.yaml

set -e

# デフォルト値
ENV=${1:-local}
# setting.yamlが同じディレクトリにあるので、それをデフォルトとして使用
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DATA_FILE=${2:-$SCRIPT_DIR/setting.yaml}
NAMESPACE="plarail"

# 環境に応じた設定
case $ENV in
  "local")
    NAMESPACE="plarail"
    ;;
  "staging")
    NAMESPACE="plarail-staging"
    ;;
  "production")
    NAMESPACE="plarail-production"
    ;;
  *)
    echo "Error: Unknown environment: $ENV"
    echo "Usage: $0 [local|staging|production] [data-file]"
    exit 1
    ;;
esac

echo "=== Seed Data Deployment ==="
echo "Environment: $ENV"
echo "Namespace: $NAMESPACE"
echo "Data file: $DATA_FILE"
echo ""

# データファイルの存在確認
if [ ! -f "$DATA_FILE" ]; then
  echo "Error: Data file not found: $DATA_FILE"
  exit 1
fi

# 既存のJobを削除（存在する場合）
echo "Cleaning up existing job..."
kubectl delete job seed-data -n $NAMESPACE --ignore-not-found=true

# ConfigMapを作成/更新
echo "Creating/updating ConfigMap..."
kubectl create configmap seed-data-config \
  --from-file=seed-data.yaml=$DATA_FILE \
  -n $NAMESPACE \
  --dry-run=client -o yaml | kubectl apply -f -

# Jobを実行
echo "Applying seed-data job..."
kubectl apply -f infra/k8s/base/seed-data/job.yaml -n $NAMESPACE

# Jobの完了を待つ
echo "Waiting for job to complete..."
kubectl wait --for=condition=complete job/seed-data -n $NAMESPACE --timeout=60s || {
  echo "Job failed or timed out. Checking logs..."
  kubectl logs job/seed-data -n $NAMESPACE
  exit 1
}

# 成功時のログを表示
echo ""
echo "=== Job completed successfully ==="
kubectl logs job/seed-data -n $NAMESPACE

# クリーンアップオプション
read -p "Do you want to delete the completed job? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
  kubectl delete job seed-data -n $NAMESPACE
  echo "Job deleted."
fi

echo ""
echo "Seed data deployment completed!"
