# Meme Coin API

A Go-based API service for meme coin operations.

## Prerequisites

Before running this application, make sure you have the following installed on your system:

- Go 1.22.3 or later
- Docker compose
- Docker
- Git

## Local Development Setup

### Option 1: Running Directly with Go

1. Clone the repository:

```bash
git clone git@github.com:nlsh710599/still-practice.git
cd meme-coin-api
```

2. Copy the environment file:

```bash
cp .env.example .env
```

3. Fill in the `.env` file with the necessary environment variables.

4. Install dependencies:

```bash
go mod download
```

5. Build and run the application:

```bash
go build -o main ./cmd/meme-coin-api
./main
```

The API will be available at `http://localhost:8080`

### Option 2: Running with Docker

1. Clone the repository:

```bash
git clone git@github.com:nlsh710599/still-practice.git
cd meme-coin-api
```

2. Before proceeding, check `docker-compose.yml` and ensure that the ports being exposed are not occupied.

3. Run the application using Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

## Project Structure

```
.
├── cmd/meme-coin-api/   # Main application entry point
├── internal/            # Internal packages
│   ├── common/         # Shared utilities
│   ├── config/         # Configuration handling
│   ├── database/       # Database operations
│   ├── route/          # API routes
│   ├── service/        # Business logic
│   └── middleware/     # Middleware for authentication and logging
├── mocks/              # Mock files for testing
├── .env                # Environment variables
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile          # Docker build file
└── README.md           # Project documentation
```

## Health Check

The application includes a health check endpoint at `/health` to verify the service status.

## Testing

### Running Tests on Your Local Machine:

If you want to run the tests on your own machine with a local PostgreSQL instance, use the following command:

```bash
PG_DSN="host=xxx user=xxx password=xxx dbname=xxx port=xxx sslmode=disable" go test ./...
```

Replace `xxx` with the correct values for your local database setup.

### Running Tests Inside the Container:

If you want to run the tests with the PostgreSQL database inside the container , use the following command:

```bash
PG_DSN="host=localhost user=postgres password=docker dbname=postgres port=5432 sslmode=disable" go test ./...
```
