Nature Remo Docker
==================
This is a Docker configuration that runs [Homebridge](https://homebridge.io/).
Helper utility for fetching/sending the IR signal is also included.

To start the Homebridge server:

```bash
$ docker-compose up -d
```

To utilize the helper:

```bash
$ docker-compose run --rm sh
~ # remo get 192.0.2.1 > data.json
~ # remo put 192.0.2.1 < data.json
```
