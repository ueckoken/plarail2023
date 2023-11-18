#include "Settings.h"

const char *ssid = "SPWH_L12_60d5d1";
const char *password = "1df7e4eabee84";

const char *HOST = "sakurajosui";

servo_item ServoSetting[] = {
    {"stop", 26, "yamashita_s1"},  // 26
    {"stop", 27, "yamashita_s2"},  // 27
    {"point", 14, "yamashita_p1"}, // 14
    {"stop", 12, "umishita_s1"}    // 12
};

switch_item SwitchSetting[] = {
    {16, "yamashita_b1", "OPEN"},
    {17, "yamashita_b1", "CLOSE"}};