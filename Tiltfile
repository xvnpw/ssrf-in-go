# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')

###
# publicapi
###

local_resource(
  'publicapi-compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/publicapi ./',
  dir="./publicapi",
  deps=['./publicapi/main.go'])

docker_build_with_restart(
  'publicapi-image',
  '.',
  entrypoint=['/app/build/publicapi'],
  dockerfile='deployments/publicapi.Dockerfile',
  only=[
    './publicapi/build',
  ],
  live_update=[
    sync('./publicapi/build', '/app/build'),
  ],
)

k8s_yaml('deployments/publicapi.yaml')
k8s_resource('publicapi', port_forwards='8001:8080',
             resource_deps=['publicapi-compile'])

### 
# backendapi
###

local_resource(
  'backendapi-compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/backendapi ./',
  dir="./backendapi",
  deps=['./backendapi/main.go'])

docker_build_with_restart(
  'backendapi-image',
  '.',
  entrypoint=['/app/build/backendapi'],
  dockerfile='deployments/backendapi.Dockerfile',
  only=[
    './backendapi/build',
  ],
  live_update=[
    sync('./backendapi/build', '/app/build'),
  ],
)

k8s_yaml('deployments/backendapi.yaml')
k8s_resource('backendapi', resource_deps=['backendapi-compile'])

### 
# imageapi
###

local_resource(
  'imageapi-compile',
  'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/imageapi ./',
  dir="./imageapi",
  deps=['./imageapi/main.go'])

docker_build_with_restart(
  'imageapi-image',
  '.',
  entrypoint=['/app/build/imageapi'],
  dockerfile='deployments/imageapi.Dockerfile',
  only=[
    './imageapi/build',
  ],
  live_update=[
    sync('./imageapi/build', '/app/build'),
  ],
)

k8s_yaml('deployments/imageapi.yaml')
k8s_resource('imageapi', resource_deps=['imageapi-compile'])