#ifndef POINTRAIL_H
#define POINTRAIL_H
#include <Arduino.h>
#include <ESP32Servo.h>

#define POINT_STRAIGHT_ANGLE 142
#define POINT_REVERSE_ANGLE 160

enum POINT_STATE
{
  POINT_STATE_NORMAL,
  POINT_STATE_REVERSE
};

class PointRail
{
  String point_id;
  POINT_STATE state;
  Servo servo;
  uint8_t pin;

public:
  PointRail();
  void attach(uint8_t pin, String point_id);
  void set_state(POINT_STATE state);
  String getId();
  POINT_STATE getState();
};

#endif