# 7sol-be-challenge

This is a Go-based backend application built based on MVC Architecture concept. It includes shorten url management and service health check.

---

## Features
- ✅ Shorten URL management
    - Minimizing
    - Reverting
---

## Running Locally without Docker
1. set these variables in a local .env file. (make sure mongo is running)

```
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=
```

2. execute following commands
```
go mod download
go run ./cmd/main.go
```

## Running with Docker
1. execute following command
```
docker-compose up --build
```
This will start:
- api (Go backend)
- postgres (Postgres database)

By default, the API runs on http://localhost:8080

## API endpoints
### HTTP
#### Public Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | /shorten | Minimize the url using the payload of original_url     |
| GET   | /:code   | Revert the code back to original url by attach the code in path param |
