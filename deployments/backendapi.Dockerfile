FROM alpine

WORKDIR /app
ADD backendapi/build build
ENTRYPOINT build/backendapi
