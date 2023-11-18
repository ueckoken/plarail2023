#ifndef SETTING_LOADER_H
#define SETTING_LOADER_H
#include <Arduino.h>
#include <ArduinoJson.h>
#include <HTTPClient.h>
#include <PubSubClient.h>
#include <LittleFS.h>
#include "Settings.h"
#include "IOManager.h"

#define FORMAT_LITTLEFS_IF_FAILED true

void getSetting(IOManager *manager);
void loadSetting(char *input, IOManager *manager);
extern bool SETTING_LOADED;

#endif