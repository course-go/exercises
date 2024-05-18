# Concurrency

This exercises focuses on the core concurrency principles common in Go like goroutines, channels, selects, and basic synchronization primitives.

## Steps

Throughout the exercise you will play around with a service that processes invoice data. The business logic does not really concern us so there isn't much of it. The processing works asynchronously with two concurrent components, the `Producer` and `Consumer` which comunicate using a channel.

First, look around the project and all of its packages. There is not much to it.

### Producer

The `Producer` component acquires new invoice ID and sends it over to the `Consumer` component.

- Implement this functionality in the `internal/producer/producer.go` file.

### Consumer

The `Consumer` component receives the invoice ID from the `Producer`, processes it, and saves it to storage.

- Implement this functionality in the `internal/consumer/consumer.go` file.

### Run It

- Wire all the dependencies and run each component in the `cmd/main.go` file.
- You can also start the `Observer` for the `Storage` which you can use to estimate the general throughput.

### Increasing Throughput

In general, the consumer and producer components are hardly ever equally fast. This results in a state, in which one of the components always has to wait for the other. We can try to match the performance of the components by scaling the slower one.

- Run multiple instance of the slower component.
    - Note that you would usually discover/verify such thresholds using benchmarks.
- You will also need to add synchronization to one of the components dependencies (either `Generator` or `Storage`).

### Graceful Shutdown

- Implement graceful shutdown for the application.
    - I.e., whenever someone will want to terminate the application, the app first flushes all its buffers and gracefuly terminates all of its componenets.
        - The `SIGINT` nad `SIGTERM` syscalls will do.
    - You will want to inspect the [signal package](https://pkg.go.dev/os/signal).
