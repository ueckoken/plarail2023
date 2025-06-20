# Seed Data Job

このJobは、MongoDBに初期データを投入するためのものです。

## 使用方法

### ローカル環境での実行

1. `infra/k8s/overlays/local/seed-data.yaml`を編集して、必要なデータを定義します。

2. k8sクラスタにデプロイします：
```bash
kubectl apply -k infra/k8s/overlays/local/
```

3. Jobの実行状況を確認します：
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
