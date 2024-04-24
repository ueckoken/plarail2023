#include "StopRail.h"

StopRail::StopRail() {}

void StopRail::attach(uint8_t pin, String stop_id)
{
  this->stop_id = stop_id;
  this->pin = pin;
}

void StopRail::set_state(STOP_STATE state)
{
  servo.attach(pin, 500, 2400);
  servo.setPeriodHertz(50);
  this->state = state;
  if (state == STOP_STATE_STOP)
  {
    servo.write(STOP_ON_ANGLE);
  }
  else
  {
    servo.write(STOP_OFF_ANGLE);
  }
  delay(200);
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