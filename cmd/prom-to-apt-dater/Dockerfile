FROM --platform=$BUILDPLATFORM golang:1.26.4-alpine@sha256:162b257ae1156df92d48e2de5abe69f657fc4961f0ea0e5376837a9ffed8160b AS builder

RUN apk add --no-cache -U git curl
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

WORKDIR /go/src/prom-to-apt-dater
COPY . /go/src/prom-to-apt-dater/

RUN --mount=type=cache,target=/go/pkg \
    go mod download -x

ARG TARGETOS
ARG TARGETARCH

RUN --mount=type=cache,target=/go/pkg \
    --mount=type=cache,target=/root/.cache/go-build \
    task generate build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM alpine:3.24@sha256:660e0827bd401543d81323d4886abbd08fda0fe3ba84337837d0b11a67251283

RUN apk add --no-cache ca-certificates mailcap

ENTRYPOINT ["/usr/bin/prom-to-apt-dater"]
CMD ["help"]

COPY --from=builder /go/src/prom-to-apt-dater/bin/prom-to-apt-dater /usr/bin/prom-to-apt-dater
WORKDIR /
