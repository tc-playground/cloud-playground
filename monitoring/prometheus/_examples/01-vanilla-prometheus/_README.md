# Docker Swarm Vanilla Prometheus

## Introduction

* Set-up and explore a vanilla `prometheus` instance (in a `Docker` container).

> NB: This example uses `Docker Stacks`.

---

## Instructions

1. Ensure `docker` is installed.

2. Initialise `docker swarm` - `make init`

3. Run `prometheus` and wait for it to start - `make start`

4. Navigate to `localhost:9090/`.

5. Play with the `built-in exporter` - This collects metrics on prometheus itself. For example golang stats.

    * List the available metrics: `curl localhost:9090/metrics`

    * Graph some of the available metrics: e.g. `go_goroutine`.

6. Tidy up - `make clean`.

---

## References

* [Docker Swarm CLI Docs](https://docs.docker.com/engine/reference/commandline)

* [Prometheus Monitoring](https://blog.alexellis.io/prometheus-monitoring/)