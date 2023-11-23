#pragma once
#ifndef TOPICROUTER_H
#define TOPICROUTER_H
#include <Arduino.h>
#include <PubSubClient.h>
#include "SettingLoader.h"
#include "Settings.h"
#include "IOManager.h"

void mqtt_handler(char *topic, byte *payload, unsigned int length, IOManager manager);

#endif