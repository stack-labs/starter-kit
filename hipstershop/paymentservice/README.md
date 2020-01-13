# Paymentservice Service

This is the Paymentservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/paymentservice --namespace=go.micro --alias=paymentservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.paymentservice
- Type: srv
- Alias: paymentservice

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
./paymentservice-srv
```

Build a docker image
```
make docker
```