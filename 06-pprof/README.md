# Profiling

The goal of this exercise is to practice generating performance profiles for Go applications and visualizing it using the [pprof](https://github.com/google/pprof) toolkit.

## Steps

### Server

We will explore profiling a HTTP server. Hence, you will need to run the server first.

```
PORT=8080 go run cmd/server/main.go
```

The server exposes a single `/data` endpoint which then does some processing and generates arbitrary data. Mind that the server can exhaust your systems resources if ran for a prolong amount of time.

### Scripts

The `scripts` directory includes scripts for generating a load on a server, so we actually have something to observe.

The first script is a basic shell script which just `curl`s the server. It can be run using `sh` like this:

```
PORT=8080 sh scripts/load.sh
```

The second is a JavaScript script that uses the [K6 load testing framework](https://k6.io) developed by Grafana Labs. To run it, you will need the `k6` executable, that can be downloaded [here](https://k6.io/docs/get-started/installation/). The script can be run like so:

```
k6 run scripts/load-k6.js
```

Both of these script will do the work. However, I suggest downloading the K6 and checking it out. It is a great tool for use cases like these.

### Profiling

You can download and inspect the profile using a single command:

```
go tool pprof "http://localhost:${PORT}/debug/pprof/profile"
```

This download the profile, saves it, opens it and runs an interactive session in which you can specify the ouput format, filters etc.

After you have downloaded the profile, you can resinspect it by using the saved file.

```
go tool pprof ~/pprof/pprof.main.alloc_objects.alloc_space.inuse_objects.inuse_space.018.pb.gz
```

You can inspect all of the exposed endpoints:

- /debug/pprof/goroutine: stack traces of all current goroutines
- /debug/pprof/heap: a sampling of memory allocations of live objects
- /debug/pprof/threadcreate: stack traces that led to the creation of new OS threads
- /debug/pprof/block: stack traces that led to blocking on synchronization primitives
- /debug/pprof/mutex: stack traces of holders of contended mutexes
- /debug/pprof/profile: cpu execution

To learn more about the specific, checkout that [pprof package](https://pkg.go.dev/runtime/pprof). Note that some of the endpoints might not return anything.
