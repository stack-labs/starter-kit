# Currencyservice Service

This is the Currencyservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/currencyservice --namespace=go.micro --alias=currencyservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.currencyservice
- Type: srv
- Alias: currencyservice

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
./currencyservice-srv
```

Build a docker image
```
make docker
```