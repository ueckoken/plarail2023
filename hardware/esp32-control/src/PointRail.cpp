#include "PointRail.h"

PointRail::PointRail() {}

void PointRail::attach(uint8_t pin, String point_id)
{
  this->point_id = point_id;
  this->pin = pin;
  servo.setPeriodHertz(50);
  servo.attach(pin, 500, 2400);
}

void PointRail::set_state(POINT_STATE state)
{
  if (state == POINT_STATE_NORMAL)
  {
    servo.write(POINT_STRAIGHT_ANGLE);
  }
  else
  {
    servo.write(POINT_REVERSE_ANGLE);
  }
  // servo.detach();
}

String PointRail::getId()
{
  return point_id;
}

POINT_STATE PointRail::getState()
{
  return state;
}