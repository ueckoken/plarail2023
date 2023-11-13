#include <Arduino.h>
#include <WiFi.h>
#include <WiFiUDP.h>
#include <ESPmDNS.h>
#include <esp_camera.h>

#define HOST "toko-minibookx"
#define PORT 3333
#define SIZE 1024

WiFiUDP udp;
IPAddress host;
uint8_t t=0;

static camera_config_t cam_cfg={// https://github.com/espressif/esp-who/blob/master/docs/en/Camera_connections.md
	.pin_pwdn=-1,.pin_reset=-1,
	.pin_xclk=4,.pin_sccb_sda=18,.pin_sccb_scl=23,

	.pin_d7=36,.pin_d6=37,.pin_d5=38,.pin_d4=39,.pin_d3=35,.pin_d2=14,.pin_d1=13,.pin_d0=34,
	.pin_vsync=5,.pin_href=27,.pin_pclk=25,

	.xclk_freq_hz=20000000,//EXPERIMENTAL: Set to 16MHz on ESP32-S2 or ESP32-S3 to enable EDMA mode
	.ledc_timer=LEDC_TIMER_0,.ledc_channel=LEDC_CHANNEL_0,

	.pixel_format=PIXFORMAT_JPEG,//YUV422,GRAYSCALE,RGB565,JPEG
	.frame_size=FRAMESIZE_SVGA,//QQVGA-UXGA, For ESP32, do not use sizes above QVGA when not JPEG. The performance of the ESP32-S series has improved a lot, but JPEG mode always gives better frame rates.

	.jpeg_quality=12, //0-63, for OV series camera sensors, lower number means higher quality
	.fb_count=2, //When jpeg mode is used, if fb_count more than one, the driver will work in continuous mode.
	.grab_mode=CAMERA_GRAB_WHEN_EMPTY//CAMERA_GRAB_LATEST. Sets when buffers should be filled
};

void setup(){
	psramInit();pinMode(21,OUTPUT);pinMode(22,OUTPUT);

	WiFi.begin();
	for(uint8_t i=0;WiFi.status()!=WL_CONNECTED;i++){
		if(i>20){
			digitalWrite(21,HIGH);
			WiFi.beginSmartConfig();while(!WiFi.smartConfigDone());
		}
		delay(500);
	}
	MDNS.begin("udp-cam");
	host=MDNS.queryHost(HOST);
	digitalWrite(21,LOW);
	esp_camera_init(&cam_cfg);
	digitalWrite(22,HIGH);
}

void loop(){
	camera_fb_t *fb=esp_camera_fb_get();
	for(uint8_t n=(fb->len+SIZE-1)/SIZE,i=0;i<n;i++){
		udp.beginPacket(host,PORT);udp.write(t);udp.write(n);udp.write(i);udp.write(fb->buf+SIZE*i,i+1==n?fb->len-SIZE*i:SIZE);udp.endPacket();
	}
	t++;
	esp_camera_fb_return(fb);
}
