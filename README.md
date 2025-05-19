# task-exporter
Receives task statuses and their duration from different external tools.

## Prerequisites
You need to have `go v1.24.2` or higher installed in order to build or run this app.

You also need to have `docker` in order to build and run the Docker image locally.

## Build
To build the task-exporter locally, run:
``` shell
go mod download && make build
```
It will create a `task-exporter` binary in the `./bin` directory.

To build the container with task-exporter, run:
``` shell
make docker_build
```

## Launch
To launch the task exporter built locally, run:
``` shell
./bin/task-exporter
```

To run a container that you've built locally, run:
``` shell
make docker_run
```

Either way, you'll have the task-exporter API running and listening on port defined by configuration (Default: `8080`)

## Local testing
There's a script that generates random events and sends them to the `task-exporter` which is running locally. You can launch it after you launched the `task-exporter` by running:
``` shell
./scripts/test_data.sh
```

After you sent enough data to the container, run this to see the metrics exposed on the `/metrics` endpoint
```
curl http://localhost:8080/metrics
```

## Configuration
Task-exporter reads configuration either from `config.env` file, or from the environment.
|Variable name|Default|Description|
|--|--|--|
|TASK_EX_PORT|`8080`|Port that task-exporter will be listening on|
|TASK_EX_ENV|`dev`|Environment that changes the way the app launches. Possible values: `dev`, `prod`|
