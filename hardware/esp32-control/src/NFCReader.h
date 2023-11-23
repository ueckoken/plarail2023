#ifndef NFCREADER_H
#define NFCREADER_H
#include <MFRC522.h>
#include <PubSubClient.h>

#define RST_PIN 9
#define SS_1_PIN 10
#define SS_2_PIN 8

class NFCReader
{
private:
  int ss_pin;
  String nfc_id;
  MFRC522 reader;
  PubSubClient *client;

public:
  NFCReader();
  void init(String nfc_id, uint8_t ss_pin, PubSubClient *client);
  void loop();
};

#endif