---
title: "About"
date: 2018-05-02T00:00:00+00:00
anchor: "about"
weight: 10
---

This command-line tool can generate automatically host configuration files for
[apt-dater][apt-dater] based on [Prometheus][prometheus] targets which get
fetched from the regular API. You can use the builtin filter mechanism powered
by [govaluate][govaluate] to limit the hosts and to gather the server group,
name and address.

[apt-dater]: https://github.com/DE-IBH/apt-dater
[prometheus]: https://prometheus.io/
[govaluate]: https://github.com/Knetic/govaluate
