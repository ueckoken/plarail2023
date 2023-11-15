#include "switch.hpp"

BlockSwitch::BlockSwitch() {}

void BlockSwitch::init(String id, String state, int pin, PubSubClient client)
{
  this->id = id;
  this->pin = pin;
  this->state = state;
  pinMode(pin, INPUT_PULLUP);
}

void BlockSwitch::loop()
{
  if (digitalRead(pin) == LOW)
  {
    if (millis() - last_switch_time < 3000)
    {
      return;
    }
    char topic[100] = "";
    sprintf(topic, "block/%s/update", id.c_str());
    client.publish(topic, state.c_str());
    last_switch_time = millis();
  }
}