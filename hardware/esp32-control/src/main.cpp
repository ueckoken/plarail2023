#include <Arduino.h>
#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>
#include <ESP32Servo.h>
#include "switch.hpp"

// WiFi
const char *ssid = "plarail-2g";
const char *password = "plarail2023";

#include "mqtt_settings.h"

// MQTT Broker
WiFiClientSecure espClient;
PubSubClient client(espClient);

typedef struct
{
  String target; // stop or point
  int pin;       // ピン番号
  String id;     // ID
} servo_item;

typedef struct
{
  int pin;      // ピン番号
  String id;    // ID
  String state; // OPEN or CLOSE
} switch_item;

#include "settings.hpp"

void callback(char *topic, byte *payload, unsigned int length);
void init_servo();
void get_init_state();

BlockSwitch switches[4];

void setup()
{
  // Set software serial baud to 115200;
  Serial.begin(9600);
  // connecting to a WiFi network
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED)
  {
    delay(500);
    Serial.print(".");
  }
  Serial.printf("\nIP address: %s\n", WiFi.localIP().toString().c_str());

  espClient.setCACert(root_ca);
  client.setServer(mqtt_broker, mqtt_port);
  client.setCallback(callback);
  while (!client.connected())
  {
    String client_id = "esp32-client-";
    client_id += String(WiFi.macAddress());
    Serial.printf("The client %s connects to the public mqtt broker\n", client_id.c_str());
    if (client.connect(client_id.c_str(), mqtt_username, mqtt_password))
    {
    }
    else
    {
      Serial.print("failed with state ");
      Serial.print(client.state());
      delay(2000);
    }
  }

  client.subscribe("stop/+/get/accepted");
  client.subscribe("point/+/get/accepted");
  client.subscribe("stop/+/delta");
  client.subscribe("point/+/delta");

  // 各種センサー類の初期化
  init_servo();
  // 初期状態の取得
  get_init_state();

  for (int i = 0; i < SWITCH_SETTING_LENGTH; i++)
  {
    switches[i].init(SwitchSetting[i].id, SwitchSetting[i].state, SwitchSetting[i].pin, client);
  }
}

void get_init_state()
{
  for (int i = 0; i < SERVO_SETTING_LENGTH; i++)
  {
    char topic[100] = "";
    if (strcmp(ServoSetting[i].target.c_str(), "stop") == 0)
    {
      sprintf(topic, "stop/%s/get", ServoSetting[i].id);
      client.publish(topic, "");
      continue;
    }
    if (strcmp(ServoSetting[i].target.c_str(), "point") == 0)
    {
      sprintf(topic, "point/%s/get", ServoSetting[i].id);
      client.publish(topic, "");
      continue;
    }
  }
}

#define GPIO_MAX 36
Servo GPIO_SERVO[GPIO_MAX];

void init_servo()
{
  for (int i = 0; i < SERVO_SETTING_LENGTH; i++)
  {
    GPIO_SERVO[ServoSetting[i].pin].setPeriodHertz(50);
    GPIO_SERVO[ServoSetting[i].pin].attach(ServoSetting[i].pin, 500, 2400);
    // if (strcmp(ServoSetting[i].target.c_str(), "stop") == 0)
    // {
    //   GPIO_SERVO[ServoSetting[i].pin].write(STOP_ON_ANGLE);
    // }
    // if (strcmp(ServoSetting[i].target.c_str(), "point") == 0)
    // {
    //   GPIO_SERVO[ServoSetting[i].pin].write(POINT_REVERSE_ANGLE);
    // }
  }
}

void change_servo_state(String target, int pin, char *state)
{
  Serial.println("change_servo_state");
  Serial.printf("target: %s, pin: %d, state: %s\n", target.c_str(), pin, state);
  if (target == "stop")
  {
    if (strcmp(state, "STOP_STATE_STOP") == 0)
    {
      Serial.println("STOP_STATE_STOP");
      GPIO_SERVO[pin].write(STOP_ON_ANGLE);
    }
    if (strcmp(state, "STOP_STATE_GO") == 0)
    {
      Serial.println("STOP_STATE_GO");
      GPIO_SERVO[pin].write(STOP_OFF_ANGLE);
    }
  }
  if (target == "point")
  {
    if (strcmp(state, "POINT_STATE_NORMAL") == 0)
    {
      GPIO_SERVO[pin].write(POINT_STRAIGHT_ANGLE);
    }
    if (strcmp(state, "POINT_STATE_REVERSE") == 0)
    {
      GPIO_SERVO[pin].write(POINT_REVERSE_ANGLE);
    }
  }
}

int get_servo_pin(String target, String id)
{
  for (int i = 0; i < SERVO_SETTING_LENGTH; i++)
  {
    if (strcmp(ServoSetting[i].target.c_str(), target.c_str()) == 0 && strcmp(ServoSetting[i].id.c_str(), id.c_str()) == 0)
    {
      return ServoSetting[i].pin;
    }
  }
  return -1;
}

void callback(char *topic, byte *payload, unsigned int length)
{
  Serial.print("Message arrived in topic: ");
  Serial.println(topic);
  Serial.print("Message:");
  for (int i = 0; i < length; i++)
  {
    Serial.print((char)payload[i]);
  }
  Serial.println();
  Serial.println("-----------------------");

  // start with stop
  if (strncmp(topic, "stop/", 5) == 0)
  {
    // get id (separator: /)
    String id = "";
    int i = 5;
    while (topic[i] != '/')
    {
      id += topic[i];
      i++;
    }
    char msg[length + 1];
    for (int i = 0; i < length; i++)
    {
      msg[i] = (char)payload[i];
    }
    msg[length] = '\0';
    change_servo_state("stop", get_servo_pin("stop", id), msg);
  }
  if (strncmp(topic, "point/", 6) == 0)
  {
    // get id (separator: /)
    String id = "";
    int i = 6;
    while (topic[i] != '/')
    {
      id += topic[i];
      i++;
    }
    char msg[length + 1];
    for (int i = 0; i < length; i++)
    {
      msg[i] = (char)payload[i];
    }
    msg[length] = '\0';
    change_servo_state("point", get_servo_pin("point", id), msg);
  }
}

void loop()
{
  client.loop();
  for (int i = 0; i < SWITCH_SETTING_LENGTH; i++)
  {
    switches[i].loop();
  }
}