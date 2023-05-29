FROM alpine

WORKDIR /app
ADD imageapi/build build
ENTRYPOINT build/imageapi
