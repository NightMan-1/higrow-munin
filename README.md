# HiGrow Plant Monitoring Sensor + Munin server

## Overview

HiGrow sensors send data to server app and then Munin collect graphs.  
Firmware basen on this - https://github.com/lucafabbri/HiGrow-Mongoose-OS-Firmware  
Server writed with GoLang and Munin plugin writed with bash.  

## How to install this app

1) install golang
2) install mos tool - https://mongoose-os.com/docs/quickstart/setup.md
3) Edit mos.yml:
	1) set higrow.deviceId and device.id
	2) set wifi.sta.ssid and wifi.sta.pass
	3) set wifi.sta.enable to true
4) edit fs/init.js - change server IP
5) compile and flash firmware
6) `go get -u github.com/kataras/iris`
7) build and run server
8) edit higrow.sh - change server IP
9) copy higrow.sh to /etc/munin/plugins folder
10) restart munin-node
11) enjoy :)


