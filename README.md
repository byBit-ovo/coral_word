# coral_world

## Run gRPC server

```
GRPC_ADDR=0.0.0.0:50051 go run .
```

## Run HTTP server (Gin)

```
HTTP_ADDR=0.0.0.0:8080 go run .
```

## Run gRPC + HTTP together

```
GRPC_ADDR=0.0.0.0:50051 HTTP_ADDR=0.0.0.0:8080 go run .
```

## etcd registration (gRPC)

Set these env vars before starting gRPC:

- `ETCD_ENDPOINTS` (comma separated, e.g. `127.0.0.1:2379`)
- `ETCD_SERVICE_NAME` (e.g. `coral_word`)

The gRPC address will be registered under `/{service}/{instance}`.
