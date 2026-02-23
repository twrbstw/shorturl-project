# 7sol-be-challenge

A Go-based backend service built using the MVC architectural pattern.
The application provides URL shortening functionality along with service health checks.

---

## Features
- ✅ URL Shortening
    - Generate a short code from an original URL
    - Revert a short code back to the original URL
- ✅ Health Monitoring
    - Liveness check
    - Readiness check
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
- postgres (PostgreSQL database)

By default, the API runs on http://localhost:8080

## API endpoints
### HTTP
#### Public Paths
| Method | Endpoint         | Description         |
|--------|------------------|---------------------|
| POST   | /shorten | Minimize the url using the payload of original_url     |
| GET   | /:code   | Revert the code back to original url by attach the code in path param |
