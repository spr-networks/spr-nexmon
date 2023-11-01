## Warning

Warning, this will overwrite your broadcom firmware on disk.
Consider first making a copy of /lib/firmware/brcm/brcmfmac43455-sdio.bin

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
tcpdump -i wlan0
```

You should see raw frames

## Changing channels
docker exec -it supernexmon /nexmon/nexutil -k36/40
docker exec -it supernexmon /nexmon/nexutil -k4/20

## Using the API

Configure the plugin path in SPR. The plugin will be reachable at http://spr.lan/plugins/nexmon/

```
curl -u admin:admin localhost/plugins/nexmon/change_channel?channel=1/20 -X PUT
```
