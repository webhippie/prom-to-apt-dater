# prom-to-apt-dater

[![Build Status](http://cloud.drone.io/api/badges/webhippie/prom-to-apt-dater/status.svg)](http://cloud.drone.io/webhippie/prom-to-apt-dater)
[![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/3395eff1e96546e7845ab0dc2173168f)](https://www.codacy.com/manual/webhippie/prom-to-apt-dater?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/prom-to-apt-dater&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/webhippie/prom-to-apt-dater?status.svg)](http://godoc.org/github.com/webhippie/prom-to-apt-dater)
[![Go Report](http://goreportcard.com/badge/github.com/webhippie/prom-to-apt-dater)](http://goreportcard.com/report/github.com/webhippie/prom-to-apt-dater)
[![](https://images.microbadger.com/badges/image/tboerger/prom-to-apt-dater.svg)](http://microbadger.com/images/tboerger/prom-to-apt-dater "Get your own image badge on microbadger.com")

This project provides a tool to generate [apt-dater](https://github.com/DE-IBH/apt-dater) host configuration from your [Prometheus](https://prometheus.io/) targets.

## Install

You can download prebuilt binaries from our [GitHub releases](https://github.com/webhippie/prom-to-apt-dater/releases), or you can use our Docker images published on [Docker Hub](https://hub.docker.com/r/webhippie/prom-to-apt-dater/tags/). If you need further guidance how to install this take a look at our [documentation](https://webhippie.github.io/prom-to-apt-dater/#getting-started).

## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/webhippie/prom-to-apt-dater.git
cd prom-to-apt-dater

make generate build

./bin/prom-to-apt-dater -h
```

## Security

If you find a security issue please contact thomas@webhippie.de first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

* [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2019 Thomas Boerger <thomas@webhippie.de>
```
