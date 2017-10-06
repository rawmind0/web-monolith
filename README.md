web-monolith
============

This image runs a web service in 8080 port, used for show monolith concept. It comes from rawmind/alpine-base.

## Build

```
docker build -t rawmind/web-monolith:<version> .
```

## Versions

- `0.1-1` [(Dockerfile)](https://github.com/rawmind0/web-monolit/blob/0.1-1/Dockerfile)


## Usage

```
docker run rawmind/web-monolith:<version> 
```

## Example

See [rancher-example][rancher-example], rancher catalog package that runs web-monolith in a cattle environment.

## Monolith modules

- /payments/ Payments module
- /inventory/ Inventory module
- /shipping/ Shipping module

[rancher-example]: https://github.com/rawmind0/web-monolith/tree/master/rancher