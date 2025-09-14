# Prom to APT-Dater

[![General Workflow](https://github.com/webhippie/prom-to-apt-dater/actions/workflows/general.yml/badge.svg)](https://github.com/webhippie/prom-to-apt-dater/actions/workflows/general.yml) [![Join the Matrix chat at https://matrix.to/#/#webhippie:matrix.org](https://img.shields.io/badge/matrix-%23webhippie-7bc9a4.svg)](https://matrix.to/#/#webhippie:matrix.org) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/3395eff1e96546e7845ab0dc2173168f)](https://app.codacy.com/gh/webhippie/prom-to-apt-dater/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Go Reference](https://pkg.go.dev/badge/github.com/webhippie/prom-to-apt-dater.svg)](https://pkg.go.dev/github.com/webhippie/prom-to-apt-dater) [![Go Report Card](https://goreportcard.com/badge/github.com/webhippie/prom-to-apt-dater)](https://goreportcard.com/report/github.com/webhippie/prom-to-apt-dater) [![GitHub Repo](https://img.shields.io/badge/github-repo-yellowgreen)](https://github.com/webhippie/prom-to-apt-dater) [![Hosted By: Cloudsmith](https://img.shields.io/badge/OSS%20hosting%20by-cloudsmith-blue?logo=cloudsmith&style=flat-square)](https://cloudsmith.com)

This project provides a tool to generate [apt-dater][aptdater] host
configurations from your [Prometheus][prometheus] targets.

## Install

You can download prebuilt binaries from the [GitHub releases][releases] or from
our [download site][downloads]. Besides that we also prepared repositories for
DEB and RPM packages which can be found at [Cloudsmith][pkgrepo]. If you prefer
to use containers you could use our images published on [GHCR][ghcr],
[Docker Hub][dockerhub] or [Quay][quay]. If you need further guidance how to
install this take a look at our [documentation][docs].

Package repository hosting is graciously provided by [Cloudsmith][cloudsmith].
Cloudsmith is the only fully hosted, cloud-native, universal package management
solution, that enables your organization to create, store and share packages in
any format, to any place, with total confidence.

## Build

If you are not familiar with [Nix][nix] it is up to you to have a working
environment for Go (>= 1.24.0) as the setup won't we covered within this guide.
Please follow the official install instructions for [Go][golang]. Beside that we
are using [go-task][gotask] to define all commands to build this project.

```console
git clone https://github.com/webhippie/prom-to-apt-dater.git
cd prom-to-apt-dater

task build
./bin/prom-to-apt-dater -h
```

If you got [Nix][nix] and [Direnv][direnv] configured you can simply execute
the following commands to get al dependencies including [go-task][gotask] and
the required runtimes installed. You are also able to directly use the process
manager of [devenv][devenv]:

```console
cat << EOF > .envrc
use flake . --impure --extra-experimental-features nix-command
EOF

direnv allow
```

## Development

To start developing on this project you have to execute only one command to
build a binary matching your platform:

```console
task build
```

After that you can simply execute the tool via `bin/prom-to-apt-dater -h`.

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
[downloads]: https://dl.webhippie.de/#prom-to-apt-dater/
[releases]: https://github.com/webhippie/prom-to-apt-dater/releases
[ghcr]: https://github.com/webhippie/prom-to-apt-dater/pkgs/container/prom-to-apt-dater
[dockerhub]: https://hub.docker.com/r/webhippie/prom-to-apt-dater/tags/
[quay]: https://quay.io/repository/webhippie/prom-to-apt-dater?tab=tags
[docs]: https://webhippie.github.io/prom-to-apt-dater/#getting-started
[nix]: https://nixos.org/
[golang]: http://golang.org/doc/install.html
[gotask]: https://taskfile.dev/installation/
[direnv]: https://direnv.net/
[devenv]: https://devenv.sh/
[pkgrepo]: https://cloudsmith.io/~webhippie/repos/general/groups/
[cloudsmith]: https://cloudsmith.com/
