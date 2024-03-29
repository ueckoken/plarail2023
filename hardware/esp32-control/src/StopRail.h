#ifndef STOPRAIL_H
#define STOPRAIL_H
#include <Arduino.h>
#include <ESP32Servo.h>

#define MAX_STOP_NUM 5
#define STOP_ON_ANGLE 180
#define STOP_OFF_ANGLE 0

enum STOP_STATE
{
  STOP_STATE_STOP,
  STOP_STATE_GO
};

class StopRail
{
  STOP_STATE state;
  Servo servo;
  uint8_t pin;

public:
  StopRail();
  void attach(uint8_t pin, String point_id);
  void set_state(STOP_STATE state);
  String getId();
  STOP_STATE getState();
  String stop_id;
};

#endif