#include "PassDetector.h"

PassDetector::PassDetector() {}

void PassDetector::init(String id, String state, int pin, PubSubClient *client)
{
  this->id = id;
  this->pin = pin;
  this->state = state;
  this->client = client;
  pinMode(pin, INPUT_PULLUP);
  this->last_switch_time = 0;
}

void PassDetector::loop()
{
  if (digitalRead(pin) == LOW)
  {
    if (millis() - last_switch_time < 3000)
    {
      return;
    }
    String topic = "pass/" + id + "/update";
    client->publish(topic.c_str(), state.c_str());
    last_switch_time = millis();
  }
}