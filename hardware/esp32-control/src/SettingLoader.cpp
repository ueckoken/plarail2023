#include "SettingLoader.h"

bool SETTING_LOADED = false;

void readFile(fs::FS &fs, const char *path, char *buf)
{
  File file = fs.open(path);
  if (!file || file.isDirectory())
  {
    Serial.println("- failed to open file for reading");
    return;
  }

  int i = 0;
  while (file.available())
  {
    *(buf + i) = file.read();
    i++;
  }
  *(buf + i) = '\0';

  file.close();
}

void loadSetting(char *input, IOManager *manager)
{
  StaticJsonDocument<384> doc;

  DeserializationError error = deserializeJson(doc, input);

  if (error)
  {
    Serial.print("deserializeJson() failed: ");
    Serial.println(error.c_str());
    return;
  }

  // 端末名
  const char* host_name = doc["name"];
  strcpy(HOST, host_name);

  // STOPS
  JsonArray stops = doc["stops"];
  for (JsonVariant v : stops)
  {
    const char *stop_id = v["stop_id"];
    int pin = v["pin"];
    Serial.printf("stop_id: %s, pin: %d\n", stop_id, pin);
    manager->addStop(pin, stop_id);
  }

  // POINTS
  JsonArray points = doc["points"];
  for (JsonVariant v : points)
  {
    const char *point_id = v["point_id"];
    int pin = v["pin"];
    Serial.printf("point_id: %s, pin: %d\n", point_id, pin);
    manager->addPoint(pin, point_id);
  }

  // DETECTORS
  JsonArray detectors = doc["detectors"];
  for (JsonVariant v : detectors)
  {
    const char *block_id = v["block_id"];
    const char *target = v["target"];
    int pin = v["pin"];
    Serial.printf("block_id: %s, target: %s, pin: %d\n", block_id, target, pin);
  }

  // NFCS
  JsonArray nfcs = doc["nfcs"];
  for (JsonVariant v : nfcs)
  {
    const char *nfc_id = v["nfc_id"];
    int pin = v["pin"];
    Serial.printf("nfc_id: %s, pin: %d\n", nfc_id, pin);
  }
  Serial.println("Setiing loaded.");
  SETTING_LOADED = true;
}

void getSetting(IOManager *manager)
{
  char json_buf[4096];
  readFile(LittleFS, "/setting.json", json_buf); // jsonファイル読み込み
  loadSetting(json_buf, manager);
}