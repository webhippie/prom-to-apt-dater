---
title: "Building"
date: 2018-05-02T00:00:00+00:00
anchor: "building"
weight: 30
---

As this project is built with Go you need to install Go first. The installation of Go is out of the scope of this document, please follow the [official documentation](https://golang.org/doc/install). After the installation of Go you need to get the sources:

{{< highlight txt >}}
go get -d github.com/webhippie/prom-to-apt-dater
cd $GOPATH/src/github.com/webhippie/prom-to-apt-dater
{{< / highlight >}}

All required tool besides Go itself are bundled or getting automatically installed within the `GOPATH`. We are using Go modules to keep the used tools consistent to manage the dependencies. All commands to build this project are part of our `Makefile`.

{{< highlight txt >}}
make generate build
{{< / highlight >}}

Finally you should have the binary within the `bin/` folder now, give it a try with `./bin/prom-to-apt-dater -h` to see all available options.
