# Productcatalogservice Service

This is the Productcatalogservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/productcatalogservice --namespace=go.micro --alias=productcatalogservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.productcatalogservice
- Type: srv
- Alias: productcatalogservice

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
./productcatalogservice-srv
```

Build a docker image
```
make docker
```