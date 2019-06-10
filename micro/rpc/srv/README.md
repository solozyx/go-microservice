# Srv Service

This is the Srv service

Generated with

```
micro new go-microservice/micro/rpc/srv --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.srv
- Type: srv
- Alias: srv

## Dependencies

Micro services depend on service discovery. The default is consul.

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
./srv-srv
```

Build a docker image
```
make docker
```