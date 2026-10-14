# rest-api

A small REST API skeleton in Go built on [httprouter](https://github.com/julienschmidt/httprouter) with [logrus](https://github.com/sirupsen/logrus) logging. The `/users` resource is wired up with all CRUD routes; the handlers are currently stubs that return placeholder responses.

## Run

```bash
go mod tidy
go run ./cmd/main
```

The server listens on port `8080` (read/write timeouts: 15 s).

## Routes

| Method | Path           | Handler             | Status |
|--------|----------------|---------------------|--------|
| GET    | `/users`       | `GetList`           | 200    |
| POST   | `/users`       | `CreateUser`        | 201    |
| GET    | `/users/:uuid` | `GetUserByUUID`     | 200    |
| PUT    | `/users/:uuid` | `UpdateUser`        | 204    |
| PATCH  | `/users/:uuid` | `PartialUpdateUser` | 204    |
| DELETE | `/users/:uuid` | `DeleteUser`        | 204    |

Example:

```bash
curl -i http://localhost:8080/users
```

## Project layout

```
cmd/main/            application entry point (router setup, HTTP server)
internal/handlers/   Handler interface implemented by every resource
internal/user/       user routes, model and storage placeholders
pkg/logging/         logrus-based logger (stdout and `logs/all.log`)
```

## Status

Early-stage learning project: persistence and real request/response handling are not implemented yet.
