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

// 接続されているStopRailとServoの設定

// target: stop, point, pin: [0-9]*, id: str

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