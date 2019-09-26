---
title: "Getting Started"
date: 2018-05-02T00:00:00+00:00
anchor: "getting-started"
weight: 20
---

## Installation

So far we are offering two different variants for the installation. You can choose between [Docker](https://www.docker.com/) or pre-built binaries which are stored on our download mirror and GitHub releases. Maybe we will also provide system packages for the major distributions later if we see the need for it.

### Docker

TBD

### Binaries

TBD

## Configuration

We provide overall three different variants of configuration. The variant based on environment variables and commandline flags are split up into global values and command-specific values.

### Envrionment variables

If you prefer to configure the service with environment variables you can see the available variables below.

#### Global

PROM_TO_APTDATER_LOG_LEVEL
: Set logging level, defaults to `info`

PROM_TO_APTDATER_LOG_COLOR
: Enable colored logging, defaults to `true`

PROM_TO_APTDATER_LOG_PRETTY
: Enable pretty logging, defaults to `true`

#### Generate

PROM_TO_APTDATER_PROMETHEUS_URL
: URL to access Prometheus, defaults to `http://localhost:9090`

PROM_TO_APTDATER_PROMETHEUS_USERNAME
: Username to access Prometheus, no default value

PROM_TO_APTDATER_PROMETHEUS_PASSWORD
: Password to access Prometheus, no default value

PROM_TO_APTDATER_OUTPUT_FILTER
: Query to filter host results, no default value

PROM_TO_APTDATER_OUTPUT_GROUP
: Attribute to define the group, defaults to `job`

PROM_TO_APTDATER_OUTPUT_NAME
: Attribute to name the host, defaults to `[__address__]`

PROM_TO_APTDATER_OUTPUT_HOST
: Attribute to access the host, defaults to `[__address__]`

PROM_TO_APTDATER_OUTPUT_FILE
: Path to generated hosts file, no default value

PROM_TO_APTDATER_OUTPUT_TEMPLATE
: Path to optional hosts template, no default value

### Commandline flags

If you prefer to configure the service with commandline flags you can see the available variables below.

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
: Username to access Prometheus, no default value

--prometheus-password
: Password to access Prometheus, no default value

--output-filter
: Query to filter host results, no default value

--output-group
: Attribute to define the group, defaults to `job`

--output-name
: Attribute to name the host, defaults to `[__address__]`

--output-host
: Attribute to access the host, defaults to `[__address__]`

--output-file
: Path to generated hosts file, no default value

--output-template
: Path to optional hosts template, no default value

### Configuration file

So far we support the file formats `JSON` and `YAML`, if you want to get a full example configuration just take a look at [our repository](https://github.com/webhippie/prom-to-apt-dater/tree/master/config), there you can always see the latest configuration format. These example configurations include all available options and the default values. The configuration file will be automatically loaded if it's placed at `/etc/prom-to-apt-dater/config.yml`, `${HOME}/.prom-to-apt-dater/config.yml` or `$(pwd)/prom-to-apt-dater/config.yml`.

## Usage

The program provides a few sub-commands on execution. The available configuration methods have already been mentioned above. Generally you can always see a formated help output if you execute the binary via `prom-to-apt-dater --help`.

### Generate

The generate command is used to generate a hosts configuration and writing it to the defined path. For further help please execute:

{{< highlight txt >}}
prom-to-apt-dater generate --help
{{< / highlight >}}
