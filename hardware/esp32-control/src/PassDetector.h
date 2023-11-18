#ifndef PASSDETECTOR_H
#define PASSDETECTOR_H

#include <Arduino.h>
#include <PubSubClient.h>

class PassDetector
{
private:
  String id;
  String state;
  int pin;
  unsigned long last_switch_time = 0;
  PubSubClient client;

public:
  PassDetector();
  void loop();
  void init(String id, String state, int pin, PubSubClient client);
};

#endif