# Example Service

This is the Example service

Generated with

```
micro new github.com/micro-in-cn/console/api --namespace=go.micro --alias=console --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.console
- Type: api
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
./console-api
```

Build a docker image
```
make docker
```
