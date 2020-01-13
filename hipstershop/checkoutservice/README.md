# Checkoutservice Service

This is the Checkoutservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/checkoutservice --namespace=go.micro --alias=checkoutservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.checkoutservice
- Type: srv
- Alias: checkoutservice

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
./checkoutservice-srv
```

Build a docker image
```
make docker
```