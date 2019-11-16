# Static Server

An extremely minimalistic and lightweight static web server, meant to be used in multi-stage Docker builds.

## Example

### Organize your Dockerfile
```dockerfile
# stage 0
FROM node:latest as builder-node

# build here your frontend source

# stage 1
FROM garsaud/static-server

COPY --from=builder-node \
    /dist \
    /www
```

### Build and execute your app
```bash
docker build -t my-application .
docker run --rm -it -p 80:80 my-application
```
