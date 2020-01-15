# Emailservice Service

This is the Emailservice service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/hipstershop/emailservice --namespace=go.micro --alias=emailservice --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.emailservice
- Type: srv
- Alias: emailservice

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
./emailservice-srv
```

Build a docker image
```
make docker
```