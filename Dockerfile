# Stage 0
FROM golang:1.13-alpine as builder-go

WORKDIR /app

COPY main.go /app/

RUN mkdir ./bin
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux \
    go build -ldflags="-s -w" -tags netgo -installsuffix netgo -o \
    ./bin/static-server

# Stage 1
FROM scratch

WORKDIR /

COPY --from=builder-go /app/bin/ /

ENTRYPOINT ["/static-server"]
