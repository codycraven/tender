# Tender

When creating microservices, the need for a simple file server often arises to serve production assets from prior build steps. Tender was designed to meet this need with no dependencies for the executable and with an extremely permissive license ([ISC](https://github.com/codycraven/tender/blob/master/LICENSE)).

The name of the project comes from purpose built boats that handle specific functions—some carry people from cruise ships to docks, some maintain buoys, some transport mail, etc.

This project utilizes the concept of tenders to provide different server functionality, the current tenders are:

- `file`

    Serves a single file at a specified route.

- `directory`

    Serves all files underneath a directory and provides a list of all files within the directory when the directory itself is requested.

- `directory no listing`

    Serves all files underneath a directory without providing a listing of all files within the directory.

## Notice

At this point in time, the API used within the Go code is absolutely prone to change. While the intent of this service is simply file serving, there may be additional tenders introduced. If you have suggestions for improving the API please feel free to open an issue or PR.

While the Go code's API may change, the configuration file format should remain backwards compatible.

## Usage

Tender requires a YAML configuration file to start. By default the configuration file is looked for in the current directory and named `config.yml`.

If you'd like to use a configuration file by a different name or in a different directory use the `--config-file` argument.

Example:

```bash
tenderserver --config-file ../some-tender-config.yml
```

## Development

Working on this project requires having [Go 1.11 or greater installed](https://golang.org/doc/install).

### Building

To build the tender server:

1. Go into the cmd/tenderserver directory

    ```bash
    cd cmd/tenderserver
    ```

1. If you cloned the repo within your $GOPATH, manually activate module mode:

    ```bash
    export GO111MODULE=on
    ```

1. Build the binary

    ```bash
    go build -mod=vendor ./...
    ```

### Dependency management

This repository makes use of [Go's modules](https://github.com/golang/go/wiki/Modules), with [vendoring](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away) for dependency management.

When a dependency is added to the codebase, `go mod vendor` should be ran.
