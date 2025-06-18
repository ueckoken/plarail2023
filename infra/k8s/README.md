# Plarail2023 Kubernetes Deployment

このディレクトリには、Plarail2023プロジェクトをKubernetesで動かすためのマニフェストが含まれています。

## 前提条件

- Kubernetesクラスター（シングルノード）が動作していること
- `kubectl`がインストールされ、クラスターに接続できること
- Dockerがインストールされていること（ローカルビルドの場合）

## ディレクトリ構造

```
infra/k8s/
├── base/                    # 基本的なKubernetesリソース
│   ├── mongodb/            # MongoDB + Mongo Express
│   ├── emqx/              # EMQXメッセージブローカー
│   ├── state-manager/     # State Managerサービス
│   ├── auto-operation/    # 自動運転サービス
│   ├── frontend/          # フロントエンドダッシュボード
│   ├── proxy/             # プロキシサービス
│   └── nginx-ingress/     # Nginx Ingressルール
├── overlays/              # 環境別の設定
│   └── local/            # ローカル開発環境用
├── build-images.sh       # Dockerイメージビルドスクリプト
└── deploy.sh            # デプロイスクリプト
```

## デプロイ手順

### 1. Dockerイメージのビルド

```bash
cd infra/k8s
./build-images.sh
```

### 2. Kubernetesへのデプロイ

```bash
./deploy.sh
```

## サービスへのアクセス

### NodePort経由（デフォルト）

- Frontend (Proxy経由): http://localhost:30031
- EMQX Dashboard: http://localhost:31808
  - ユーザー名: admin
  - パスワード: password
- Mongo Express: http://localhost:30081
  - ユーザー名: admin
  - パスワード: password
- MQTT: localhost:31883

### Nginx Ingress経由（オプション）

Nginx Ingress Controllerがインストールされている場合：

```bash
# Nginx Ingress Controllerのインストール
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/cloud/deploy.yaml
```

インストール後、以下のURLでアクセス可能：

- Frontend: http://localhost/
- State Manager API: http://localhost/api/
- EMQX Dashboard: http://localhost/emqx/
- Mongo Express: http://localhost/mongo-express/

## 管理コマンド

### ステータス確認

```bash
# すべてのリソースを確認
kubectl -n plarail get all

# Podのログを確認
kubectl -n plarail logs -f deployment/state-manager
```

### 削除

```bash
# すべてのリソースを削除
kubectl delete -k overlays/local/

# Namespaceごと削除
kubectl delete namespace plarail
```

## トラブルシューティング

### Podが起動しない場合

```bash
# Podの詳細を確認
kubectl -n plarail describe pod <pod-name>

# イベントを確認
kubectl -n plarail get events --sort-by='.lastTimestamp'
```

### PVCがPendingの場合

ローカル環境でStorageClassが設定されていない場合は、以下を実行：

```bash
# デフォルトのStorageClassを作成
kubectl apply -f - <<EOF
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
EOF
```

## カスタマイズ

### イメージタグの変更

`overlays/local/kustomization.yaml`でイメージタグを変更できます：

```yaml
images:
  - name: plarail2023/state-manager
    newTag: v1.0.0  # 任意のタグに変更
```

### 環境変数の追加

各サービスのdeployment.yamlで環境変数を追加できます。

### リソース制限の設定

本番環境では、各コンテナにリソース制限を設定することを推奨します：

```yaml
resources:
  requests:
    memory: "256Mi"
    cpu: "250m"
  limits:
    memory: "512Mi"
    cpu: "500m"
