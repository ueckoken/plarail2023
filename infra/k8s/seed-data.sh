#!/bin/bash

# Seed Data実行スクリプト
# 使用方法: ./seed-data.sh [データファイル]
# 例: ./seed-data.sh overlays/local/setting.yaml

set -e

# デフォルト値
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
NAMESPACE="plarail"

# データファイルのパス（引数で指定されない場合はデフォルトを使用）
DATA_FILE=${1:-$SCRIPT_DIR/overlays/local/setting.yaml}

echo "=== Seed Data Deployment ==="
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
kubectl apply -f $SCRIPT_DIR/base/seed-data/job.yaml -n $NAMESPACE

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
