# GoATAK - free ATAK/CivTAK server & web-based client

This is Golang implementation of ATAK server/CoT router aimed to test some ideas about CoT message routing.

binary builds can be downloaded
from [releases page](https://github.com/kdudkov/goatak/releases)

![Alt text](client.png?raw=true "Title")

## Web-based client features

* v1 (XML) and v2 (protobuf) CoT protocol support
* SSL connection support, tested with [FreeTakServer](https://github.com/FreeTAKTeam/FreeTakServer)
  and [Argustak](https://argustak.com/)
* web-ui, ideal for big screen situation awareness center usage
* unit track - your target unit is always in the center of map
* RedX tool - to measure distance and bearing
* Digital Pointer - send DP position to all other contacts
* Add and edit units on map

## GoATAK server features

* v1 (XML) and v2 (protobuf) CoT protocol support
* certificate enrollment (v1 and v2) support

[Server fast start](https://github.com/kdudkov/goatak/wiki/Setting-up-server)

## Web client setup

1. Download latest binary build
   from [releases page](https://github.com/kdudkov/goatak/releases)
1. Unzip it to local directory
1. edit `goatak_client.yml` (default values are for community server).
1. run `webclient`
1. open [http://localhost:8080](http://localhost:8080) in your browser

You can use as many config files as you want and run with specific config with `webclient -config <your_config.yml>`

### Web client config examples

simple config to connect to [Argustak](https://argustak.com/) cloud based TAK server:

```yaml
---
server_address: argustak.com:4444:ssl
web_port: 8080
me:
   callsign: username
   uid: auto
   type: a-f-G-U-C
   team: Blue
   role: Team Member
   lat: 0
   lon: 0
ssl:
   cert: username.p12
   password: password
```

## Libraries used

* [Leaflet](https://leafletjs.com/)
* [Milsymbol](https://github.com/spatialillusions/milsymbol)

[![CI](https://github.com/kdudkov/goatak/actions/workflows/main.yml/badge.svg?branch=master)](https://github.com/kdudkov/goatak/actions/workflows/main.yml)

[By me a beer 🍺](https://buymeacoffee.com/kdudkov)
