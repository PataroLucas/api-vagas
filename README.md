# api-vagas

RESTful API for managing job openings, built with Go, Gin, GORM, and SQLite.

## Features

- **CRUD operations** for job openings (Create, Read, Update, Delete, List)
- **SQLite persistence** with auto-migration via GORM
- **Structured logging** with debug, info, warn, and error levels
- **Clean architecture** separating config, handlers, router, and schemas

## Tech Stack

- [Go](https://go.dev/) 1.26
- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://gorm.io/) — ORM with SQLite driver
- [SQLite](https://www.sqlite.org/) — Embedded database

## Getting Started

### Prerequisites

- Go 1.26+

### Installation

```bash
git clone https://github.com/PataroLucas/api-vagas.git
cd api-vagas
go mod tidy
go run main.go
```

The server starts on `http://localhost:8080`.

## API Endpoints

All endpoints are prefixed with `/api/v1`.

| Method   | Endpoint       | Description          |
|----------|----------------|----------------------|
| `GET`    | `/opening`     | Show a job opening   |
| `POST`   | `/opening`     | Create a job opening |
| `PUT`    | `/opening`     | Update a job opening |
| `DELETE` | `/opening`     | Delete a job opening |
| `GET`    | `/openings`    | List all openings    |

## Project Structure

```
api-vagas/
├── config/          # Configuration, logger, and SQLite initialization
│   ├── config.go
│   ├── logger.go
│   └── sqlite.go
├── db/              # SQLite database file (auto-created)
├── handler/         # HTTP request handlers
│   ├── handler.go
│   ├── request.go
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
├── go.mod
└── go.sum
```

## License

[MIT](LICENSE)
