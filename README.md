# ssrf-in-go

Presentation about SSRF in Go from BSides Conference

## Requirements

Install:

- [ctlptl](https://github.com/tilt-dev/ctlptl)
- [kind](https://kind.sigs.k8s.io)
- [tilt](https://tilt.dev)

## Setup

Create kubernetes cluster using kind:

```
ctlptl create cluster kind --registry=ctlptl-registry
```

## Run

```
tilt up
```

```
curl -s http://localhost:8001/debug\?url\=http://backendapi/internal
```
