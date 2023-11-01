#!/bin/bash

cd /nexmon
cp brcmfmac43455-sdio.bin /lib/firmware/cypress/cyfmac43455-sdio-standard.bin

rmmod brcmfmac_wcc
rmmod brcmfmac

insmod brcmfmac.ko

sleep 1

iw phy `iw dev wlan0 info | awk '/wiphy/ {printf "phy" $2}'` interface add mon0 type monitor

echo [+] Loaded

cd /
/nexmon_plugin
