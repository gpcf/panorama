# Panorama

Interactive isometric maps for Minetest

<a href="http://map.lord-server.ru">
    <small><i>Click here to see a live instance</i></small>
    <img src="https://user-images.githubusercontent.com/4698994/163820087-6473cbc4-b790-4e6d-9130-aedb5bf1eddf.png"></img>
</a>

## Installation

*Note: Panorama started as our in-house mapper, and installation is
non-trivial as a result. If you're not comfortable with complicated
setups, check out [mapserver] instead!*

### Prerequisites

- PostgreSQL backend for your world
- Several gigabytes of disk space for tiles
- A decent CPU and about a gigabyte of RAM, depending on workload
- [`nodes_dump`][nodes_dump] mod installed

### Using Docker (recommended)

This is an easier option, especially if you already use a Docker-based setup for
your server. There are pre-built [Docker images][docker-image] that you can use,
or you can build it yourself using provided Dockerfile.

Here's an example `docker-compose.yml` to get you started:

```yml
version: "3"
services:
  panorama:
    image: ghcr.io/lord-server/panorama:latest
    ports:
      - "33333:33333"
    volumes:
      - "/path/to/minetest/worlds/my-world:/var/lib/panorama/world"
      - "/path/to/minetest/games/minetest_game:/var/lib/panorama/game"
      - "/path/to/config/dir:/etc/panorama"
      - "/path/to/tiles:/var/lib/panorama/tiles"
    command: ["--serve", "--fullrender"]
```

### Building manually

Coming soon! Follow commands in [Dockerfile](https://github.com/lord-server/panorama/blob/master/Dockerfile) in the meantime.

## License

MIT

[instance]: http://map.lord-server.ru/
[mapserver]: https://github.com/minetest-mapserver/mapserver
[docker-image]: https://github.com/lord-server/panorama/pkgs/container/panorama
[nodes_dump]: https://github.com/lord-server/nodes_dump
