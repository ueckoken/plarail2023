#ifndef SETTINGS_H
#define SETTINGS_H
#include <Arduino.h>

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

// 接続されているStopRailとServoの設定

// target: stop, point, pin: [0-9]*, id: str

#define SERVO_SETTING_LENGTH 4
#define SWITCH_SETTING_LENGTH 2

extern servo_item ServoSetting[];

extern switch_item SwitchSetting[];

// WiFi
// const char *ssid = "plarail-2g";
// const char *password = "plarail2023";
extern const char *ssid;
extern const char *password;

extern const char *HOST;

#endif