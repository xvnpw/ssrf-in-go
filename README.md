# ssrf-in-go

Presentation about SSRF in Go from BSides Conference

## Install

Create kubernetes cluster using kind

```
ctlptl create cluster kind --registry=ctlptl-registry
```

## Run

```
tilt up
```

```
curl -s \
http://localhost:8001/debug\?url\=http://backendapi/internal
```
