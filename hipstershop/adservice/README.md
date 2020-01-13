# Adservice Service

This is the Adservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/adservice --namespace=go.micro --alias=adservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.adservice
- Type: srv
- Alias: adservice

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./adservice-srv
```

Build a docker image
```
make docker
```