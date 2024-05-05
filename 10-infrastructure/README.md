# Infrastructure

This exercise is split into two parts:
- The goal of the first task is to set-up a Caddy server instance as a reverse proxy for a multiple services.
- The second part includes more of a demo in which you will discover the Google Cloud Platform together with your tutor.

## Caddy

### Ping-Pong

The set-up uses [an existing repository](https://github.com/course-go/ping-pong) with a Go HTTP server. The server exposes a single endpoint `/ping`. Feel free to pull it, run it and try it out.

1. Checkout the `compose` set-up.
    - It contains multiple instances of the `ping-pong` API.
2. Add [caddy](https://hub.docker.com/_/caddy) to the compose set-up.
    - You will also need to mount a volume to propagate the server config created in the next step.
3. Create a `Caddyfile` and set-up a reverse proxy for the `ping-pong` instances.
    - You can find more about the `Caddyfile` directives in the [documentation](https://caddyserver.com/docs/caddyfile).
    - Make sure the Caddy server balances load between the instances.
        - You can use the round-robin stategy.
4. Test it out.
    - When quering the `/ping` endpoint the responses should cycle between the instance numbers.
    - The instances should also be accessible only via the proxy.

```sh
curl localhost:80/ping
```

## GCP

Make sure you have you have singed up and are ready to use the Google Cloud Platform Console. The demo mainly consists of interactions using the UI so this exercise contains no written guide.
