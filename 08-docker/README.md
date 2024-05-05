# Docker

The goal of this exercise is to practice containarizing a HTTP API. The API is a implementation of the Todo app introduced in the previous exercise. If you want to use your own solution instead, you can certainly do so.

## Steps

### Dockerfile

The first step is to create a custom image.

1. Create a new `Dockerfile` in the root of the project.
2. Set-up the Dockerfile with the necessary steps.
    - There is [golang](https://hub.docker.com/_/golang) image on Docker Hub you can use as the starting point.
    - You will want to copy all source code to the image and build it.
    - Finally, execute the outputted binary as the last step.
3. [Optional] Use multi-stage build to lower the image size.
    - You can use the [alpine](https://hub.docker.com/_/alpine) image as the runner.
        - Note that you might want to also use the **golang:alpine** image for the builder, otherwise the final binary may not work.
4. Test the container image.
    - Build it and run it.
    - You will have to expose the right port for the API.
    - Revisit the previous exercise for hints how to query the API.

### Docker Compose

As running `docker run` repetitively with the right arguments is tidious, we will want to create a compose file.

1. Create a new `compose.yaml` in the root of the project.
2. Set-up a single `api` service.
    - It will reuse the `Dockerfile` at the build step.
3. Run the compose and test out the API again.
