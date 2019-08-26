# Prometheus

## Introduction

* `prometheus` is **time series database** and **monitoring system**. 

* `prometheus` server **scrapes** metrics from **instrumented jobs**, either directly or via an intermediary **push gateway** for short-lived jobs.

* `prometheus` supports concepts of `exporters` - a service or daemon to collect metrics for consumption.

* `prometheus` takes time-series data and treats it as a data source for **generating alerts**.


---

## Features

* A multi-dimensional data model with `time series` data identified by `metric name` and `key/value pairs`.

* `PromQL` a flexible query language to leverage this dimensionality.
 
* No reliance on distributed storage; single server nodes are autonomous.
 
* Time series collection happens via a `pull model over HTTP`.
 
* `Pushing time series` is supported via an intermediary gateway.

* targets are discovered via service discovery or static configuration.
 
* Multiple modes of `graphing` and `dashboarding` support.


---

## Architecture

![Prometheus Architecture](./prometheus-architecture.png)

---

## Tutorials

* [Instrument a Golang Server in 5 minute](https://medium.com/@gsisimogang/instrumenting-golang-server-in-5-min-c1c32489add3)

* [Instrument Golang Service - RED Method](https://dev.to/dzeban/instrumenting-a-go-service-for-prometheus-khp)

* [Prometheus Monitoring](https://blog.alexellis.io/prometheus-monitoring/)

* [Monitoring Go Applications with Prometheus](https://scot.coffee/2018/12/monitoring-go-applications-with-prometheus/)

* __Prometheus Series__

    * [01 - Metrics and Labels](https://blog.pvincent.io/2017/12/prometheus-blog-series-part-1-metrics-and-labels/)

    * [02 - Metric Types](https://blog.pvincent.io/2017/12/prometheus-blog-series-part-2-metric-types/)

    * [03 - Exposing and Collecting Metrics](https://blog.pvincent.io/2017/12/prometheus-blog-series-part-3-exposing-and-collecting-metrics/)

    * [04 - Instrumenting Code](https://blog.pvincent.io/2017/12/prometheus-blog-series-part-4-instrumenting-code-in-go-and-java/)

    * [05 - Alerting Rules](https://blog.pvincent.io/2017/12/prometheus-blog-series-part-5-alerting-rules/)


---

## References

* [Home](https://prometheus.io/)

* [Github](https://github.com/prometheus)

* [Documentation](https://prometheus.io/docs/introduction/overview/)

    * [Data Model](https://prometheus.io/docs/concepts/data_model/)

    * [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/)

        * [Examples](https://prometheus.io/docs/prometheus/latest/querying/examples/)

        * [RE2 syntax](https://github.com/google/re2/wiki/Syntax)

    * [Instrument a Golang Application](https://prometheus.io/docs/guides/go-application/)

    * [Alert Configuration](https://prometheus.io/docs/alerting/configuration/)

