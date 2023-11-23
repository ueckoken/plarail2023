#include "PointRail.h"

PointRail::PointRail() {}

void PointRail::attach(uint8_t pin, String point_id)
{
  point_id = point_id;
  pin = pin;
  servo.setPeriodHertz(50);
}

void PointRail::set_state(POINT_STATE state)
{
  state = state;
  servo.attach(pin, 500, 2400);
  if (state == POINT_STATE_NORMAL)
  {
    servo.write(POINT_STRAIGHT_ANGLE);
  }
  else
  {
    servo.write(POINT_REVERSE_ANGLE);
  }
  servo.detach();
}

String PointRail::getId()
{
  return point_id;
}

POINT_STATE PointRail::getState()
{
  return state;
}