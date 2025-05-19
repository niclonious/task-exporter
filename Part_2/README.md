# Part 2

## Prerequisites

- You've successfully built the task-exporter Docker container using the default values in this repo's makefile.
```
make docker_build
```
- You have docker-compose installed
- You need to `cd` to this directory

## Running Part 2
Run the bash script provided in this directory:
```
./run.sh
```
This will copy prometheus config file to `/tmp` directory to mount it inside a prometheus container which will be launched together with task-exporter by running `compose up`.

After that you can run the script provided in this repository to feed test data to the task-exporter.

> The prometheus UI should be available at http://localhost:9090

## Cleanup
After you're done running Part 2, you can stop the containers and then run `cleanup.sh` provided in this directory do remove prometheus config from your system and remove both containers.
```
./cleanup.sh
```

## Prometheus query:
Here's the query that you can use in prometheus to see the highest task duration per tool:
``` promql
max by (__name__, tool) (task_duration)
```