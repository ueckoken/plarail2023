#include <Arduino.h>
#include <PubSubClient.h>
#include <WiFi.h>
#include <WiFiClientSecure.h>
#include <ESP32Servo.h>

// WiFi
const char *ssid = "plarail-2g";
const char *password = "plarail2023";

// MQTT Broker
const char *mqtt_broker = "p390e24a.ala.us-east-1.emqxsl.com"; // broker address
const char *topic = "stop/#";                                  // define topic
const char *mqtt_username = "test";                            // username for authentication
const char *mqtt_password = "password";                        // password for authentication
const int mqtt_port = 8883;                                    // port of MQTT over TLS
const char *root_ca =
    "-----BEGIN CERTIFICATE-----\n"
    "MIIEqjCCA5KgAwIBAgIQAnmsRYvBskWr+YBTzSybsTANBgkqhkiG9w0BAQsFADBh\n"
    "MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\n"
    "d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\n"
    "QTAeFw0xNzExMjcxMjQ2MTBaFw0yNzExMjcxMjQ2MTBaMG4xCzAJBgNVBAYTAlVT\n"
    "MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\n"
    "b20xLTArBgNVBAMTJEVuY3J5cHRpb24gRXZlcnl3aGVyZSBEViBUTFMgQ0EgLSBH\n"
    "MTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALPeP6wkab41dyQh6mKc\n"
    "oHqt3jRIxW5MDvf9QyiOR7VfFwK656es0UFiIb74N9pRntzF1UgYzDGu3ppZVMdo\n"
    "lbxhm6dWS9OK/lFehKNT0OYI9aqk6F+U7cA6jxSC+iDBPXwdF4rs3KRyp3aQn6pj\n"
    "pp1yr7IB6Y4zv72Ee/PlZ/6rK6InC6WpK0nPVOYR7n9iDuPe1E4IxUMBH/T33+3h\n"
    "yuH3dvfgiWUOUkjdpMbyxX+XNle5uEIiyBsi4IvbcTCh8ruifCIi5mDXkZrnMT8n\n"
    "wfYCV6v6kDdXkbgGRLKsR4pucbJtbKqIkUGxuZI2t7pfewKRc5nWecvDBZf3+p1M\n"
    "pA8CAwEAAaOCAU8wggFLMB0GA1UdDgQWBBRVdE+yck/1YLpQ0dfmUVyaAYca1zAf\n"
    "BgNVHSMEGDAWgBQD3lA1VtFMu2bwo+IbG8OXsj3RVTAOBgNVHQ8BAf8EBAMCAYYw\n"
    "HQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMBIGA1UdEwEB/wQIMAYBAf8C\n"
    "AQAwNAYIKwYBBQUHAQEEKDAmMCQGCCsGAQUFBzABhhhodHRwOi8vb2NzcC5kaWdp\n"
    "Y2VydC5jb20wQgYDVR0fBDswOTA3oDWgM4YxaHR0cDovL2NybDMuZGlnaWNlcnQu\n"
    "Y29tL0RpZ2lDZXJ0R2xvYmFsUm9vdENBLmNybDBMBgNVHSAERTBDMDcGCWCGSAGG\n"
    "/WwBAjAqMCgGCCsGAQUFBwIBFhxodHRwczovL3d3dy5kaWdpY2VydC5jb20vQ1BT\n"
    "MAgGBmeBDAECATANBgkqhkiG9w0BAQsFAAOCAQEAK3Gp6/aGq7aBZsxf/oQ+TD/B\n"
    "SwW3AU4ETK+GQf2kFzYZkby5SFrHdPomunx2HBzViUchGoofGgg7gHW0W3MlQAXW\n"
    "M0r5LUvStcr82QDWYNPaUy4taCQmyaJ+VB+6wxHstSigOlSNF2a6vg4rgexixeiV\n"
    "4YSB03Yqp2t3TeZHM9ESfkus74nQyW7pRGezj+TC44xCagCQQOzzNmzEAP2SnCrJ\n"
    "sNE2DpRVMnL8J6xBRdjmOsC3N6cQuKuRXbzByVBjCqAA8t1L0I+9wXJerLPyErjy\n"
    "rMKWaBFLmfK/AHNF4ZihwPGOc7w6UHczBZXH5RFzJNnww+WnKuTPI0HfnVH8lg==\n"
    "-----END CERTIFICATE-----\n";

WiFiClientSecure espClient;
PubSubClient client(espClient);

// 接続されているStopRailとServoの設定

// target: stop, point, pin: [0-9]*, id: str

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

/*
  Stations:
    yamashita_station
    umishita_station
  StopRail:
    yamashita_s1
    yamashita_s2
    umishita_s1
  PointRail:
    yamashita_p1
  PassDetector:
    yamashita_d1
    umishita_d1
*/

#define STOP_ON_ANGLE 45
#define STOP_OFF_ANGLE 5

#define POINT_STRAIGHT_ANGLE 165
#define POINT_REVERSE_ANGLE 180

#define SERVO_SETTING_LENGTH 4
#define SWITCH_SETTING_LENGTH 2

servo_item ServoSetting[] = {
    {"stop", 26, "yamashita_s1"},  // 26
    {"stop", 27, "yamashita_s2"},  // 27
    {"point", 14, "yamashita_p1"}, // 14
    {"stop", 12, "umishita_s1"}    // 12
};

switch_item SwitchSetting[] = {
    {16, "yamashita_b1", "OPEN"},
    {17, "yamashita_b1", "CLOSE"}};

void callback(char *topic, byte *payload, unsigned int length);
void init_switch();
void init_servo();
void get_init_state();

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
  init_switch();
  init_servo();
  // 初期状態の取得
  get_init_state();
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

void init_switch()
{
  for (int i = 0; i < SWITCH_SETTING_LENGTH; i++)
  {
    pinMode(SwitchSetting[i].pin, INPUT_PULLUP);
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

unsigned long LAST_SWITCH_TIME[SWITCH_SETTING_LENGTH] = {0};

void switch_check()
{
  for (int i = 0; i < SWITCH_SETTING_LENGTH; i++)
  {
    if (digitalRead(SwitchSetting[i].pin) == LOW)
    {
      if (millis() - LAST_SWITCH_TIME[i] < 3000)
      {
        continue;
      }
      char topic[100] = "";
      sprintf(topic, "block/%s/update", SwitchSetting[i].id.c_str());
      client.publish(topic, SwitchSetting[i].state.c_str());
      LAST_SWITCH_TIME[i] = millis();
    }
  }
}

void loop()
{
  client.loop();
  switch_check();
}