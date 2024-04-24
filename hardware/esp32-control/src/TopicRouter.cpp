#include "TopicRouter.h"

void mqtt_handler(char *topic, byte *payload, unsigned int length, IOManager *manager)
{
  /*
    受信したTopicに応じて処理を分岐
  */

  // start with stop
  if (strncmp(topic, "stop/", 5) == 0)
  {
    // get id (separator: /)
    String id = "";
    int i = 5;
    while (topic[i] != '/')
    {
      id += topic[i];
      i++;
    }
    char msg[length + 1];
    for (int i = 0; i < length; i++)
    {
      msg[i] = (char)payload[i];
    }
    msg[length] = '\0';

    if (strcmp(msg, "STOP_STATE_GO") == 0)
    {
      manager->setStopState(id, STOP_STATE_GO);
    }
    else if (strcmp(msg, "STOP_STATE_STOP") == 0)
    {
      manager->setStopState(id, STOP_STATE_STOP);
    }

    return;
  }

  if (strncmp(topic, "point/", 6) == 0)
  {
    // get id (separator: /)
    String id = "";
    int i = 6;
    while (topic[i] != '/')
    {
      id += topic[i];
      i++;
    }
    char msg[length + 1];
    for (int i = 0; i < length; i++)
    {
      msg[i] = (char)payload[i];
    }
    msg[length] = '\0';

    if (strcmp(msg, "POINT_STATE_NORMAL") == 0)
    {
      manager->setPointState(id, POINT_STATE_NORMAL);
    }
    else if (strcmp(msg, "POINT_STATE_REVERSE") == 0)
    {
      manager->setPointState(id, POINT_STATE_REVERSE);
    }

    return;
  }
}