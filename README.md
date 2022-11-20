# QRGen

> NOTE: Not production ready. Use at your own risk.

1. `go install github.com/bradj/qrgen`
1. `qrgen wifi -s [ssid] -p [password]`

## Generate code for WiFi network
```
Generates a QR code for the spcified wifi ssid and password

Usage:
  qrgen wifi [flags]

Flags:
  -h, --help              help for wifi
  -o, --output string     Specify the location of the generated png (default "wifi-qr.png")
  -p, --password string   The password of your wifi network
  -s, --ssid string       The ssid of your wifi network
```

## Used Libraries

* [qr - By AlexEidt](https://github.com/AlexEidt/qr) 👏
