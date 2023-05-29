FROM alpine

WORKDIR /app
ADD publicapi/build build
ENTRYPOINT build/publicapi
