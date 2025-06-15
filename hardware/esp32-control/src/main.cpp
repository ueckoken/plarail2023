#include <Arduino.h>
#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiClient.h>
#include <ESP32Servo.h>
#include <ArduinoOTA.h>
#include <LittleFS.h>
#include <ESPmDNS.h>

// Include the header file of the library
#include "PassDetector.h"
#include "PointRail.h"
#include "StopRail.h"
#include "MQTTSetting.h"
#include "Settings.h"
#include "TopicRouter.h"
#include "SettingLoader.h"
#include "IOManager.h"

#define DEBUG

#define SPI_SCK 0
#define SPI_MISO 2
#define SPI_MOSI 1

// MQTT Broker
WiFiClient espClient;
PubSubClient client(espClient);
IOManager manager(&client);

void callback(char *topic, byte *payload, unsigned int length);
void reconnect(); // 再接続関数を追加

void setup()
{
  Serial.begin(115200);

  // PubSubClientのバッファサイズを増やす（デフォルト256から1024へ）
  client.setBufferSize(1024);

  // WiFiに接続
  Serial.println("-----Welcome to Plarail IoT System!-----");
  WiFi.begin(ssid, password);
  Serial.print("MAC Address: ");
  Serial.println(WiFi.macAddress());
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.printf("\nIP address: %s\n", WiFi.localIP().toString().c_str());

  // LittleFS初期化
  if (!LittleFS.begin(FORMAT_LITTLEFS_IF_FAILED))
  {
    Serial.println("LittleFS Mount Failed");
    return;
  }

  // // SPIを初期化
  // SPI.begin(SPI_SCK, SPI_MISO, SPI_MOSI);

  // 設定ファイルを読み込む
  getSetting(&manager);

  if (!SETTING_LOADED)
  {
    Serial.println("Setting File Not Found");
    return;
  }

  // MDNSを開始
  MDNS.begin(HOST);

  // espClient.setCACert(root_ca);
  client.setServer(mqtt_broker, mqtt_port);
  client.setCallback(callback);
  client.setKeepAlive(60);     // KeepAliveを60秒に設定
  client.setSocketTimeout(30); // ソケットタイムアウトを30秒に設定

  reconnect(); // 初回接続を再接続関数で行う

  // 初期状態を取得
  manager.getInitialState();
}

// 再接続関数を追加
void reconnect()
{
  while (!client.connected())
  {
    String client_id = "esp32-client-";
    client_id += String(HOST);
    Serial.printf("The client %s connects to the public mqtt broker\n", client_id.c_str());

    if (client.connect(client_id.c_str(), mqtt_username, mqtt_password))
    {
      Serial.println("MQTT connected");
      // 再接続時にTopicを再購読
      client.subscribe("stop/+/get/accepted");
      client.subscribe("point/+/get/accepted");
      client.subscribe("stop/+/delta");
      client.subscribe("point/+/delta");
      Serial.println("Topics resubscribed after reconnection");
    }
    else
    {
      Serial.print("failed with state ");
      int state = client.state();
      Serial.print(state);
      switch (state)
      {
      case -4:
        Serial.println(" - MQTT_CONNECTION_TIMEOUT");
        break;
      case -3:
        Serial.println(" - MQTT_CONNECTION_LOST");
        break;
      case -2:
        Serial.println(" - MQTT_CONNECT_FAILED");
        break;
      case -1:
        Serial.println(" - MQTT_DISCONNECTED");
        break;
      case 1:
        Serial.println(" - MQTT_CONNECT_BAD_PROTOCOL");
        break;
      case 2:
        Serial.println(" - MQTT_CONNECT_BAD_CLIENT_ID");
        break;
      case 3:
        Serial.println(" - MQTT_CONNECT_UNAVAILABLE");
        break;
      case 4:
        Serial.println(" - MQTT_CONNECT_BAD_CREDENTIALS");
        break;
      case 5:
        Serial.println(" - MQTT_CONNECT_UNAUTHORIZED");
        break;
      default:
        Serial.println(" - Unknown error");
      }
      Serial.println("Retrying in 5 seconds...");
      delay(5000);
    }
  }
}

void callback(char *topic, byte *payload, unsigned int length)
{
#ifdef DEBUG
  Serial.print("Message arrived in topic: ");
  Serial.println(topic);
  Serial.print("Message:");
  for (int i = 0; i < length; i++)
  {
    Serial.print((char)payload[i]);
  }
  Serial.println();
  Serial.println("-----------------------");
#endif

  mqtt_handler(topic, payload, length, &manager);
}

void loop()
{
  // WiFi接続チェック
  if (WiFi.status() != WL_CONNECTED)
  {
    Serial.println("WiFi disconnected. Reconnecting...");
    WiFi.reconnect();
    delay(1000);
    return;
  }

  // MQTT接続チェック
  if (!client.connected())
  {
    Serial.println("MQTT disconnected. Reconnecting...");
    reconnect();
  }

  manager.loop();
  client.loop();
}
