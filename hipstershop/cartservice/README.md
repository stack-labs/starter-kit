# Cartservice Service

This is the Cartservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/cartservice --namespace=go.micro --alias=cartservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.cartservice
- Type: srv
- Alias: cartservice

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
./cartservice-srv
```

Build a docker image
```
make docker
```