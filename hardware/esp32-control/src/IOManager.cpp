#include "IOManager.h"

StopRail STOP_LIST[MAX_STOP_NUM];

IOManager::IOManager(PubSubClient *client)
{
  this->client = client;
  POINT_LIST_INDEX = 0;
  STOP_LIST_INDEX = 0;
  DETECTOR_LIST_INDEX = 0;
  NFC_LIST_INDEX = 0;
}

void IOManager::addPoint(uint8_t pin, String point_id)
{
  POINT_LIST[POINT_LIST_INDEX] = PointRail();
  POINT_LIST[POINT_LIST_INDEX].attach(pin, point_id);
  POINT_LIST_INDEX++;
}

void IOManager::addStop(uint8_t pin, String stop_id)
{
  STOP_LIST[STOP_LIST_INDEX].attach(pin, stop_id);
  STOP_LIST_INDEX++;
}

void IOManager::addDetector(uint8_t pin, String block_id, String target)
{
  DETECTOR_LIST[DETECTOR_LIST_INDEX]
      .init(block_id, target, pin, client);
  DETECTOR_LIST_INDEX++;
}

void IOManager::addNfc(uint8_t pin, String nfc_id)
{
  NFC_LIST[NFC_LIST_INDEX].init(nfc_id, pin, client);
  NFC_LIST_INDEX++;
}

void IOManager::setPointState(String point_id, POINT_STATE state)
{
  for (int i = 0; i < POINT_LIST_INDEX; i++)
  {
    if (POINT_LIST[i].getId() == point_id)
    {
      POINT_LIST[i].set_state(state);
      return;
    }
  }
}

void IOManager::setStopState(String stop_id, STOP_STATE state)
{
  for (int i = 0; i < STOP_LIST_INDEX; i++)
  {
    Serial.println(STOP_LIST[i].stop_id);
    if (STOP_LIST[i].stop_id == stop_id)
    {
      Serial.println(STOP_LIST[i].getId());
      Serial.println(stop_id);
      STOP_LIST[i].set_state(state);
      return;
    }
  }
}

void IOManager::getInitialState()
{
  for (int i = 0; i < POINT_LIST_INDEX; i++)
  {
    String topic = "point/" + POINT_LIST[i].getId() + "/get";
    this->client->publish(topic.c_str(), "get");
  }
  for (int i = 0; i < STOP_LIST_INDEX; i++)
  {
    String topic = "stop/" + STOP_LIST[i].stop_id + "/get";
    this->client->publish(topic.c_str(), "get");
  }
}

void IOManager::loop()
{
  for (int i = 0; i < DETECTOR_LIST_INDEX; i++)
  {
    DETECTOR_LIST[i].loop();
  }
  for (int i = 0; i < NFC_LIST_INDEX; i++)
  {
    NFC_LIST[i].loop();
  }
}