# Prom to APT-Dater

[![Current Tag](https://img.shields.io/github/v/tag/webhippie/prom-to-apt-dater?sort=semver)](https://github.com/webhippie/prom-to-apt-dater) [![Build Status](https://github.com/webhippie/prom-to-apt-dater/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/prom-to-apt-dater/actions) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/prom-to-apt-dater.svg)](https://pkg.go.dev/github.com/webhippie/prom-to-apt-dater) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/prom-to-apt-dater)](https://goreportcard.com/report/github.com/webhippie/prom-to-apt-dater) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/3395eff1e96546e7845ab0dc2173168f)](https://www.codacy.com/gh/webhippie/prom-to-apt-dater/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=webhippie/prom-to-apt-dater&amp;utm_campaign=Badge_Grade)

This project provides a tool to generate [apt-dater][aptdater] host
configurations from your [Prometheus][prometheus] targets.


## Install

You can download prebuilt binaries from our [GitHub releases][releases], or you
can use our Docker images published on [Docker Hub][dockerhub] or [Quay][quay].
If you need further guidance how to install this take a look at our
[documentation][docs].

## Development

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions][golang]. This project requires
Go >= v1.17, at least that's the version we are using.

```console
git clone https://github.com/webhippie/prom-to-apt-dater.git
cd prom-to-apt-dater

make generate build

./bin/prom-to-apt-dater -h
```

## Security

If you find a security issue please contact
[thomas@webhippie.de](mailto:thomas@webhippie.de) first.

## Contributing

Fork -> Patch -> Push -> Pull Request

## Authors

-   [Thomas Boerger](https://github.com/tboerger)

## License

Apache-2.0

## Copyright

```console
Copyright (c) 2019 Thomas Boerger <thomas@webhippie.de>
```

[aptdater]: https://github.com/DE-IBH/apt-dater
[prometheus]: https://prometheus.io/
[releases]: https://github.com/webhippie/prom-to-apt-dater/releases
[dockerhub]: https://hub.docker.com/r/webhippie/prom-to-apt-dater/tags/
[quay]: https://quay.io/repository/webhippie/prom-to-apt-dater?tab=tags
[docs]: https://webhippie.github.io/prom-to-apt-dater/#getting-started
[golang]: http://golang.org/doc/install.html
