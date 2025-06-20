# Seed Data Job

このJobは、MongoDBに初期データを投入するためのものです。

## 使用方法

### スクリプトを使用した実行（推奨）

`infra/k8s/seed-data.sh`スクリプトを使用すると、簡単にseed-dataを実行できます：

```bash
# デフォルト（local環境、infra/k8s/setting.yaml）
./infra/k8s/seed-data.sh

# 環境とデータファイルを指定
./infra/k8s/seed-data.sh local backend/onetime/seed-data/data/chofufes-2023.yaml

# staging環境で実行
./infra/k8s/seed-data.sh staging path/to/your/data.yaml
```

デフォルトでは、`infra/k8s/setting.yaml`が使用されます。

スクリプトは以下を自動的に行います：
- 既存のJobのクリーンアップ
- ConfigMapの作成/更新
- Jobの実行と完了待機
- ログの表示

### 手動での実行

1. k8sクラスタにデプロイします（`infra/k8s/setting.yaml`を使用）：
```bash
kubectl apply -k infra/k8s/overlays/local/
```

2. Jobの実行状況を確認します：
```bash
kubectl get jobs -n plarail
kubectl logs -n plarail job/seed-data
```

### 本番環境での実行

1. 実際のデータファイルからConfigMapを作成します：
```bash
kubectl create configmap seed-data-config \
  --from-file=seed-data.yaml=path/to/your/data.yaml \
  -n your-namespace
```

2. Jobを実行します：
```bash
kubectl apply -f infra/k8s/base/seed-data/job.yaml -n your-namespace
```

## 環境変数

- `MONGODB_URI`: MongoDBの接続URI
- `SEED_DATA_FILE`: 読み込むYAMLファイルのパス（デフォルト: `./data/mfk-2024.yaml`）

## YAMLファイルの形式

```yaml
stop_rails:
  - "rail_id_1"
  - "rail_id_2"
point_rails:
  - "point_id_1"
  - "point_id_2"
blocks:
  - "block_id_1"
  - "block_id_2"
