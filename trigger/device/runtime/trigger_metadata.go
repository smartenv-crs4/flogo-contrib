package device

var jsonMetadata = `{
  "name": "tibco-device",
  "version": "0.0.1",
  "type": "device",
  "title": "Device Trigger",
  "description": "Simple Device Trigger",
  "settings":[
    {
      "name": "mqtt_server",
      "type": "string",
      "required" : true
    },
    {
      "name": "mqtt_user",
      "type": "string"
    },
    {
      "name": "mqtt_password",
      "type": "string"
    },
    {
      "name": "device:name",
      "type": "string",
      "required" : true
    },
    {
      "name": "device:ssid",
      "type": "string",
      "required" : true
    },
    {
      "name": "device:wifi_password",
      "type": "string",
      "required" : true
    },
    {
      "name": "device:board",
      "type": "string",
      "required" : true,
      "allowed" : ["feather_m0_wifi"]
    }
  ],
  "endpoint": {
    "settings": [
      {
        "name": "device:pin",
        "type": "string",
        "required" : true
      },
      {
        "name": "device:condition",
        "type": "string",
        "required" : true
      },
      {
        "name": "device:response_pin",
        "type": "string"
      }
    ]
  }
}`