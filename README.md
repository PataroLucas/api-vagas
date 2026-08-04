# api-vagas

RESTful API for managing job openings, built with Go, Gin, GORM, and SQLite.

## Features

- **CRUD operations** for job openings (Create, Read, Update, Delete, List)
- **SQLite persistence** with auto-migration via GORM
- **Structured logging** with debug, info, warn, and error levels
- **Clean architecture** separating config, handlers, router, and schemas
- **Swagger/OpenAPI documentation** served at `/swagger/*any`
- **Makefile** with common tasks (run, build, test, docs, clean)

## Tech Stack

- [Go](https://go.dev/) 1.26
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://gorm.io/) — ORM with SQLite driver
- [SQLite](https://www.sqlite.org/) — Embedded database
- [Swagger](https://github.com/swaggo/swag) — OpenAPI docs generation
- [gin-swagger](https://github.com/swaggo/gin-swagger) — Swagger UI

## Getting Started

### Prerequisites

- Go 1.26+
- [swag](https://github.com/swaggo/swag) CLI for generating Swagger docs:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Installation

```bash
git clone https://github.com/PataroLucas/api-vagas.git
cd api-vagas
go mod tidy
make run
```

Or use the default target, which regenerates the docs before starting:

```bash
make run-with-docs
```

The server starts on `http://localhost:8080`.

### Available Make Targets

| Target          | Description                                   |
|-----------------|-----------------------------------------------|
| `make run`      | Run the server                                |
| `make run-with-docs` | Regenerate docs, then run the server (default) |
| `make build`    | Build the binary                              |
| `make test`     | Run the tests                                 |
| `make docs`     | Regenerate Swagger docs                       |
| `make clean`    | Remove the binary and generated docs          |

## Swagger

Interactive API documentation is available at `http://localhost:8080/swagger/index.html`. The docs are generated from annotations in the handlers using `swag init` (or `make docs`).

## API Endpoints

All endpoints are prefixed with `/api/v1`.

| Method   | Endpoint       | Parameters         | Description          |
|----------|----------------|--------------------|----------------------|
| `GET`    | `/opening`     | `?id=<id>`         | Show a job opening   |
| `POST`   | `/opening`     | JSON body          | Create a job opening |
| `PUT`    | `/opening`     | `?id=<id>` + body  | Update a job opening |
| `DELETE` | `/opening`     | `?id=<id>`         | Delete a job opening |
| `GET`    | `/openings`    | —                  | List all openings    |

### Opening fields

- `role` (string)
- `company` (string)
- `location` (string)
- `remote` (bool)
- `link` (string)
- `salary` (int64)

Responses are wrapped in `{ "message": string, "data": <opening | opening[]> }`. Errors return `{ "message": string, "errorCode": int }`.

## Project Structure

```
api-vagas/
├── config/          # Configuration, logger, and SQLite initialization
│   ├── config.go
│   ├── logger.go
│   └── sqlite.go
├── db/              # SQLite database file (auto-created, gitignored)
├── docs/            # Generated Swagger docs (docs.go, swagger.json, swagger.yaml)
├── handler/         # HTTP request handlers
│   ├── handler.go
│   ├── request.go
│   ├── response.go
│   ├── createOpening.go
│   ├── deleteOpening.go
│   ├── listOpenings.go
│   ├── showOpening.go
│   └── updateOpening.go
├── router/          # Router setup and route definitions
│   ├── router.go
│   └── routes.go
├── schemas/         # Data models and response types
│   └── opening.go
├── main.go          # Application entry point
├── makefile         # Common tasks (run, build, test, docs, clean)
├── go.mod
└── go.sum
```

## License

[MIT](LICENSE)
