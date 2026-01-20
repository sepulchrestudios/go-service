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

## Troubleshooting

### Seeing TLS errors about unrecognized certificate signing authority?

Copy the relevant CA certificate into the `build/certs` directory and make sure it has a `.crt` extension.

The `Dockerfile` steps are configured to auto-copy certs in that location and update the CA bundle in the image when building.

### Unsure how to retrieve a CA certificate?

You can use the `openssl` binary from the command line to get its X.509-formatted data.

See [this GitLab link with instructions](https://forum.gitlab.com/t/gitlab-ci-cd-runner-registration-tls-failed-to-verify-certificate-x509-certificate-signed-by-unknown-authority/102007/2) to get the relevant data.

The steps are fundamentally as follows (massaged from that link):

1. `openssl s_client -showcerts -connect <hostname>:443 </dev/null 2>/dev/null | openssl x509 -outform PEM > <filename>.crt`
2. `cp <filename>.crt build/certs/`

Replace `<hostname>` with the relevant hostname from which you wish to retrieve the certificate.

Replace `<filename>` with the relevant filename you wish to generate.

Rebuild and restart the Docker container and you should be all set!