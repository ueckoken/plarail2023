# Device Manager

MQTT Brokerを経由してESP32に対して命令を出力、受け取るためのサービスです。

やることとしては、基本的にProtobufで定義されたメッセージをMQTT BrokerにPublishすることです。