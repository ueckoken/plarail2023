#ifndef IOMANAGER_H
#define IOMANAGER_H
#include <Arduino.h>
#include <PassDetector.h>
#include <PointRail.h>
#include <StopRail.h>

#define MAX_POINT_NUM 5
#define MAX_STOP_NUM 5
#define MAX_DETECTOR_NUM 5
#define MAX_NFC_NUM 5

class IOManager
{
public:
  IOManager(PubSubClient client);
  PubSubClient client;
  void addStop(uint8_t pin, String stop_id);
  void addPoint(uint8_t pin, String point_id);

  void setStopState(String stop_id, STOP_STATE state);
  void setPointState(String point_id, POINT_STATE state);
  void addDetector(uint8_t pin, String block_id, String target);
  // addNfc(uint8_t pin, String block_id);

  void loop();

private:
  uint8_t POINT_LIST_INDEX;
  uint8_t STOP_LIST_INDEX;
  uint8_t DETECTOR_LIST_INDEX;
  // uint8_t NFC_LIST_INDEX;
  PointRail POINT_LIST[MAX_POINT_NUM];
  StopRail STOP_LIST[MAX_STOP_NUM];
  PassDetector DETECTOR_LIST[MAX_DETECTOR_NUM];
  // NfcReader NFC_LIST[MAX_NFC_NUM];
};

#endif