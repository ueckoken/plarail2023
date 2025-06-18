# State Manager

State Managerは、プラレールシステムの状態を管理するサービスです。

## 修正された問題点

### 1. エラーハンドリングの改善
- `log.Fatal`の使用を削除し、適切なエラーハンドリングに変更
- MQTTハンドラーのStart()メソッドの戻り値処理を修正
- リカバリーミドルウェアを有効化してパニック時のクラッシュを防止

### 2. リソース管理の改善
- データベースカーソルの適切なクローズ処理を追加
- データベース操作にタイムアウトを設定
- シャットダウン時のデータベース接続クローズ処理を追加

### 3. 設定の柔軟性向上
- 設定ファイルのパスを環境変数で指定可能に
- TLS設定のエラーハンドリングを改善

## 必要な環境変数

以下の環境変数を`.env`ファイルまたはシステム環境変数に設定してください：

```bash
# MongoDB接続設定
MONGODB_URI=mongodb://localhost:27017

# MQTT接続設定
MQTT_BROKER_ADDR=ssl://your-mqtt-broker:8883
MQTT_USERNAME=your-username
MQTT_PASSWORD=your-password
MQTT_CLIENT_ID=state-manager-001

# アプリケーション設定
APP_ENV=dev  # dev, test, prod

# オプション：設定ファイルのパス（デフォルト: ../settings/esp）
SETTINGS_PATH=/path/to/settings/esp
```

## TLS証明書

MQTTブローカーへの接続にはTLS証明書が必要です。`emqxsl-ca.pem`ファイルを実行ディレクトリに配置してください。

## 起動方法

```bash
# 開発環境
go run cmd/main.go

# 本番環境
APP_ENV=prod ./state-manager
```

## トラブルシューティング

### サービスが起動しない場合
1. 環境変数が正しく設定されているか確認
2. MongoDBが起動しているか確認
3. MQTTブローカーが起動しているか確認
4. TLS証明書ファイルが存在するか確認

### サービスが落ちる場合
1. ログを確認して具体的なエラーメッセージを特定
2. MQTTトピックの形式が正しいか確認（`target/id/method`の形式）
3. データベース接続が安定しているか確認
