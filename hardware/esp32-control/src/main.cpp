#include <Arduino.h>
#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>
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

// MQTT Broker
WiFiClientSecure espClient;
PubSubClient client(espClient);

IOManager manager(client);

void callback(char *topic, byte *payload, unsigned int length);

void setup()
{
  Serial.begin(115200);
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

  MDNS.begin(HOST);

  espClient.setCACert(root_ca);
  client.setServer(mqtt_broker, mqtt_port);
  client.setCallback(callback);
  while (!client.connected())
  {
    String client_id = "esp32-client-";
    client_id += String(HOST);
    Serial.printf("The client %s connects to the public mqtt broker\n", client_id.c_str());
    if (!client.connect(client_id.c_str(), mqtt_username, mqtt_password))
    {
      Serial.print("failed with state ");
      Serial.print(client.state());
      delay(2000);
    }
  }

  // Subscribe Topics
  client.subscribe("stop/+/get/accepted");
  client.subscribe("point/+/get/accepted");
  client.subscribe("stop/+/delta");
  client.subscribe("point/+/delta");

  // LittleFS初期化
  if (!LittleFS.begin(FORMAT_LITTLEFS_IF_FAILED))
  {
    Serial.println("LittleFS Mount Failed");
    return;
  }
  getSetting(&manager);
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

  mqtt_handler(topic, payload, length, manager);
}

void loop()
{
  manager.loop();
  client.loop();
}