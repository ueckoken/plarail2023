#include "NFCReader.h"

NFCReader::NFCReader()
{
}

void NFCReader::init(String nfc_id, uint8_t ss_pin, PubSubClient *client)
{
    this->ss_pin = ss_pin;
    this->nfc_id = nfc_id;
    this->client = client;
    reader.PCD_Init(ss_pin, RST_PIN);
}

void NFCReader::loop()
{
    Serial.println(reader.PICC_IsNewCardPresent());
    if (!reader.PICC_IsNewCardPresent())
    {
        return;
    }
    if (!reader.PICC_ReadCardSerial())
    {
        return;
    }
    String uid = "";
    for (byte i = 0; i < reader.uid.size; i++)
    {
        uid += String(reader.uid.uidByte[i] < 0x10 ? "0" : "");
        uid += String(reader.uid.uidByte[i], HEX);
    }
    Serial.print("NFC ID: ");
    Serial.print(nfc_id);
    Serial.print("UID: ");
    Serial.println(uid);
    reader.PICC_HaltA();
    reader.PCD_StopCrypto1();

    // Publish Topic
    String topic = "nfc/" + nfc_id + "/update";
    client->publish(topic.c_str(), uid.c_str());
}