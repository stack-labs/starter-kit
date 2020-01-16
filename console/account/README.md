# Example Service

This is the Example service

Generated with

```
micro new github.com/micro-in-cn/starter-kit/console/account --namespace=go.micro --alias=account --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.console
- Type: srv
- Alias: console

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./console-srv
```

Build a docker image
```
make docker
```
