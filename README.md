# Sepulchre Studios - Psychokinesis (PK)

## About

Developed by [@matthewfritz](https://github.com/matthewfritz).

_Psychokinesis_ (PK) is the framework upon which all other Go-based services for Sepulchre Studios are intended to be built.

An example `main.go` file is provided to show a "liveness" server running and give you an idea of how to implement your own service startup sequence.

The example service runs within the provided development containers, configured within the `docker-compose.yml` file.

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
