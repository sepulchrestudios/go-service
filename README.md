# Sepulchre Studios - Go Service

## About

Developed by [@matthewfritz](https://github.com/matthewfritz).

This is the core layer that is intended to serve as a module dependency for our dedicated services, written in Go.

An example `main.go` file is provided to show a "liveness" server running and give you an idea of how to implement your own service startup sequence.

This example service runs within the provided development containers, configured within the `docker-compose.yml` file.

## Configuration

### Generating the Protocol Buffer Files

```
make proto
```

**NOTE:** you should **only** need to do this if you have actively changed the source `proto/*.proto` files in some way.

### Copying Environment Configuration Files

```
make copy-env
```

**NOTE:** this is **not** required to be done prior to running `make start`, as it is handled automatically.

## Development Containers

### Building the Go Server

```
make build
```

### Starting the Containers

```
make start
```

### Building the Go Server and Starting the Containers

```
make dev
```

### Stopping the Containers

```
make stop
```
