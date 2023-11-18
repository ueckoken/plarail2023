#include "StopRail.h"

StopRail::StopRail() {}

void StopRail::attach(uint8_t pin, String point_id)
{
  point_id = point_id;
  pin = pin;
  servo.setPeriodHertz(50);
}

void StopRail::set_state(STOP_STATE state)
{
  state = state;
  servo.attach(pin, 500, 2400);
  if (state == STOP_STATE_STOP)
  {
    servo.write(STOP_ON_ANGLE);
  }
  else
  {
    servo.write(STOP_OFF_ANGLE);
  }
  servo.detach();
}

String StopRail::getId()
{
  return stop_id;
}

STOP_STATE StopRail::getState()
{
  return state;
}