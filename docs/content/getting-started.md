---
title: "Getting Started"
date: 2022-05-04T00:00:00+00:00
anchor: "getting-started"
weight: 20
---

## Installation

So far we are offering only a few different variants for the installation. You
can choose between [Docker][docker] or pre-built binaries which are stored on
our download mirror and GitHub releases. Maybe we will also provide system
packages for the major distributions later if we see the need for it.

### Docker

Generally we are offering the images through
[quay.io/webhippie/prom-to-apt-dater][quay] and
[webhippie/prom-to-apt-dater][dockerhub], so feel free to choose one of the
providers. Maybe we will come up with Kustomize manifests or some Helm chart.

### Binaries

Simply download a binary matching your operating system and your architecture
from our [downloads][downloads] or the GitHub releases and place it within your
path like `/usr/local/bin` if you are using macOS or Linux.

## Configuration

We provide overall three different variants of configuration. The variant based
on environment variables and commandline flags are split up into global values
and command-specific values.

### Envrionment variables

If you prefer to configure the service with environment variables you can see
the available variables below.

#### Global

PROM_TO_APT_DATER_LOG_LEVEL
: Set logging level, defaults to `info`

PROM_TO_APT_DATER_LOG_COLOR
: Enable colored logging, defaults to `true`

PROM_TO_APT_DATER_LOG_PRETTY
: Enable pretty logging, defaults to `true`

### Generate

PROM_TO_APT_DATER_LOG_LEVEL
: Set logging level, defaults to `info`

PROM_TO_APT_DATER_PROMETHEUS_URL
: URL to access Prometheus, defaults to `http://localhost:9090`

PROM_TO_APT_DATER_PROMETHEUS_USERNAME
: Username to access Prometheus

PROM_TO_APT_DATER_PROMETHEUS_PASSWORD
: Password to access Prometheus

PROM_TO_APT_DATER_OUTPUT_FILTER
: Query to filter host results, defaults to `1 == 1`

PROM_TO_APT_DATER_OUTPUT_GROUP
: Attribute to define the group, defaults to `job`

PROM_TO_APT_DATER_OUTPUT_NAME
: Attribute to name the host, defaults to `[__address__]`

PROM_TO_APT_DATER_OUTPUT_USER
: Attribute to detect the user, defaults to `static('root')`

PROM_TO_APT_DATER_OUTPUT_HOST
: Attribute to access the host, defaults to `[__address__]`

PROM_TO_APT_DATER_OUTPUT_PORT
: Attribute to detect the port, defaults to `static('22')`

PROM_TO_APT_DATER_OUTPUT_FILE
: Path to generated hosts file

PROM_TO_APT_DATER_OUTPUT_TEMPLATE
: Path to optional hosts template

### Commandline flags

If you prefer to configure the service with commandline flags you can see the
available variables below.

#### Global

--log-level
: Set logging level, defaults to `info`

--log-color
: Enable colored logging, defaults to `true`

--log-pretty
: Enable pretty logging, defaults to `true`

#### Generate

--prometheus-url
: URL to access Prometheus, defaults to `http://localhost:9090`

--prometheus-username
: Username to access Prometheus

--prometheus-password
: Password to access Prometheus

--output-filter
: Query to filter host results, defaults to `1 == 1`

--output-group
: Attribute to define the group, defaults to `job`

--output-name
: Attribute to name the host, defaults to `[__address__]`

--output-user
: Attribute to detect the user, defaults to `static('root')`

--output-host
: Attribute to access the host, defaults to `[__address__]`

--output-port
: Attribute to detect the port, defaults to `static('22')`

--output-file
: Path to generated hosts file

--output-template
: Path to optional hosts template

### Configuration file

So far we support multiple file formats like `json` or `yaml`, if you want to
get a full example configuration just take a look at [our repository][repo],
there you can always see the latest configuration format. These example configs
include all available options and the default values. The configuration file
will be automatically loaded if it's placed at
`/etc/prom-to-apt-dater/config.yml`, `${HOME}/.prom-to-apt-dater/config.yml` or
`$(pwd)/prom-to-apt-dater/config.yml`.

## Usage

The program provides a few sub-commands on execution. The available config
methods have already been mentioned above. Generally you can always see a
formated help output if you execute the binary similar to something like
 `prom-to-apt-dater --help`.

[docker]: https://www.docker.com/
[quay]: https://quay.io/repository/webhippie/prom-to-apt-dater
[dockerhub]: https://hub.docker.com/r/webhippie/prom-to-apt-dater
[downloads]: https://dl.webhippie.de/
[repo]: https://github.com/webhippie/prom-to-apt-dater/tree/master/config
