# Prometheus Instrumented Golang Service

## Introduction

* Set-up and explore a vanilla `prometheus` instance (in a `Docker` container).

> NB: This example uses `pure Docker`.

---

## Instructions

1. Ensure `docker` is installed.

2. Ensure the simple Go `service` is initialised and has it's dependencies.

    ```bash
    cd service
    export GO111MODULE=on
    # go mod init service
    go mod vendor -v
    ```

3. Build and test the simple service

    ```bash 
    go build -o service
    ./service
    ```

4. 