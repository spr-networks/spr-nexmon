# About 

This SPR plugin loads seemoo lab's nexmon, which enables monitor mode on the built-in wifi radio. without nexmon this is not otherwise posible on the raspberry pi 4b.

## Warning

Warning, this will overwrite your broadcom firmware on disk.
Consider first making a copy of /lib/firmware/brcm/brcmfmac43455-sdio.bin

## Versions 

This expects the SPR pi image with a 6.2 Ubuntu kernel. 

## Building
```
docker compose build
```

## Running

```export SUPERDIR=/home/spr/super
docker compose up -d
```

## Testing

```
# tcpdump -i wlan0 ...

tcpdump: verbose output suppressed, use -v[v]... for full protocol decode
listening on wlan0, link-type IEEE802_11_RADIO (802.11 plus radiotap header), snapshot length 262144 bytes
22:50:27.005540 1876482302us tsft 1.0 Mb/s 2412 MHz 11b -68dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.046106 1876522917us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.107930 1876584711us tsft 1.0 Mb/s 2412 MHz 11b -70dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.148500 1876625317us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.210323 1876687100us tsft 1.0 Mb/s 2412 MHz 11b -67dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.250927 1876727725us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.353387 1876830134us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.415150 1876891900us tsft 1.0 Mb/s 2412 MHz 11b -67dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.455780 1876932515us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.517581 1876994317us tsft 1.0 Mb/s 2412 MHz 11b -68dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.619973 1877096718us tsft 1.0 Mb/s 2412 MHz 11b -68dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.660499 1877137311us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.722352 1877199099us tsft 1.0 Mb/s 2412 MHz 11b -68dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.762911 1877239713us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.824758 1877301507us tsft 1.0 Mb/s 2412 MHz 11b -67dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.865300 1877342110us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.927152 1877403898us tsft 1.0 Mb/s 2412 MHz 11b -69dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:27.967757 1877444502us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.029491 1877506279us tsft 1.0 Mb/s 2412 MHz 11b -68dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.070097 1877546905us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.131952 1877608687us tsft 1.0 Mb/s 2412 MHz 11b -69dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.172498 1877649302us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.267684 1877745972us tsft 1.0 Mb/s 2412 MHz 11b -63dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.268706 1877747115us tsft 1.0 Mb/s 2412 MHz 11b -62dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.269562 1877747974us tsft 1.0 Mb/s 2412 MHz 11b -63dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.274840 1877751689us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.377343 1877854092us tsft 1.0 Mb/s 2412 MHz 11b -46dBm signal 0dBm noise Beacon (wifi-2.4) [1.0* 2.0* 5.5* 11.0* 6.0 9.0 12.0 18.0 Mbit] ESS CH: 1, PRIVACY
22:50:28.393455 1877871861us tsft 1.0 Mb/s 2412 MHz 11b -60dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.394594 1877873004us tsft 1.0 Mb/s 2412 MHz 11b -60dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.395453 1877873863us tsft 1.0 Mb/s 2412 MHz 11b -61dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
22:50:28.396315 1877874722us tsft 1.0 Mb/s 2412 MHz 11b -61dBm signal 0dBm noise Probe Request (wifiap) [1.0 2.0 5.5 6.0 9.0 11.0 12.0 18.0 Mbit]
^C

```

You should see raw frames

## Changing channels
docker exec -it supernexmon /nexmon/nexutil -k36/40
docker exec -it supernexmon /nexmon/nexutil -k4/20

## Using the API

Configure the plugin path in SPR. The plugin will be reachable at http://spr.lan/plugins/nexmon/

![image](https://github.com/spr-networks/spr-nexmon/assets/37549748/79fffc7a-4897-48ce-8b20-c824911cdb68)

```
curl -u admin:admin localhost/plugins/nexmon/change_channel?channel=1/20 -X PUT
```

