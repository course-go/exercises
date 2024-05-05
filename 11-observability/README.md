# Observability

Like the last exercise, this one is also split into two parts:
- The goal of the first task is to set-up, instrument and inspect Prometheus metrics.
- The second part, again, includes more of a demo in which you will inspect an OpenTelemetry Demo together with your tutor.

## Metrics

### Prometheus

All services are already set-up for you in the `compose` file. All you need to do is to instrument the application using the Prometheus client.

1. Add [Prometheus client](https://github.com/prometheus/client_golang) to the dependencies.
2. Instrument the Todo APP.
    - Create a metric for each controller.
        - E.g. total number of request calls for each controller with HTTP response code as label.
        - Follow [this guide](https://prometheus.io/docs/practices/naming/) when naming your metrics and labels.
    - You can use `promauto` to simplify the set-up.
3. Expose the metrics.
    - You can use a custom endpoint, but you will have to specify it in the Prometheus config file.
4. Change prometheus config in `/configs/prometheus` directory to scrape the metrics.
5. Inspect the metrics.
    - You can start with just `curl`ing the metrics endpoint.
    - After that, move to the Prometheus UI and try to create a simple graph.

### Grafana

1. Move on to the Grafana UI.
2. Create a dashboard for the Todo service.
3. Create a time-series graph for the controllers endpoints.
    - You can also create graphs for the Go runtime metrics.

## OpenTelemetry

Clone the OpenTelemetry Demo and run it. The tutor will guide you through the set-up and all the tools it uses.
- [OpenTelemetry Demo](https://github.com/open-telemetry/opentelemetry-demo)
- [OpenTelemetry Demo Documentation](https://opentelemetry.io/docs/demo/)
