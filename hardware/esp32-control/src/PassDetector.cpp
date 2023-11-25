#include "PassDetector.h"

PassDetector::PassDetector() {}

void PassDetector::init(String id, String state, int pin, PubSubClient *client)
{
  this->id = id;
  this->pin = pin;
  this->state = state;
  pinMode(pin, INPUT_PULLUP);
}

void PassDetector::loop()
{
  if (digitalRead(pin) == LOW)
  {
    if (millis() - last_switch_time < 3000)
    {
      return;
    }
    String topic = "block/" + id + "/update";
    client->publish(topic.c_str(), state.c_str());
    last_switch_time = millis();
  }
}