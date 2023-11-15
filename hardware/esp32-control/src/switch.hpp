#include <Arduino.h>
#include <PubSubClient.h>

class BlockSwitch
{
private:
  String id;
  String state;
  int pin;
  unsigned long last_switch_time = 0;
  PubSubClient client;

public:
  BlockSwitch();
  void loop();
  void init(String id, String state, int pin, PubSubClient client);
};