# Project Setup Guide

This guide provides detailed instructions on how to set up and run the Go Gin Doc project locally.

## Prerequisites

Ensure you have the following installed on your system:

1.  **Go**: Version 1.23 or higher. [Download Go](https://go.dev/dl/)
2.  **PostgreSQL**: A running PostgreSQL instance. [Download PostgreSQL](https://www.postgresql.org/download/)
3.  **Make**: (Optional) For running convenience commands defined in `Makefile`.
4.  **golang-migrate**: For handling database migrations.
    - Install via CLI:
      ```bash
      # macOS (Homebrew)
      brew install golang-migrate

      # Linux (Debian/Ubuntu)
      curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
      sudo mv migrate /usr/bin/migrate

      # Windows (Scoop)
      scoop install migrate
      ```
    - Or via Go install:
      ```bash
      go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
      ```

## Installation Steps

### 1. Clone the Repository

```bash
git clone https://github.com/Walon-Foundation/go-gin-doc.git
cd go-gin-doc
```

### 2. Install Dependencies

Download the Go module dependencies:

```bash
go mod download
```

### 3. Environment Configuration

Create a `.env` file in the root directory. You can copy the structure below:

```env
# Server Configuration
PORT=5000

# Authentication
JWT_SECRET=your_secure_random_string

# Database Configuration
# Format: postgres://username:password@host:port/database_name?sslmode=disable
DB_URL=postgres://walon:password@localhost:5432/app_db?sslmode=disable
```

> **Note**: Replace the values in `DB_URL` with your actual PostgreSQL credentials.

### 4. Database Setup

1.  **Create the Database**:
    Log in to your PostgreSQL instance and create the database defined in your `DB_URL` (e.g., `app_db`).
    ```sql
    CREATE DATABASE app_db;
    ```

2.  **Run Migrations**:
    Apply the database schema using the `make` command:
    ```bash
    make migrate-up
    ```
    Or manually using `migrate`:
    ```bash
    migrate -path cmd/db/migrations -database "postgres://user:password@localhost:5432/app_db?sslmode=disable" up
    ```

### 5. Generate Swagger Docs (Optional)

If you modify the API comments, regenerate the Swagger documentation:

> **Learning Tip**: Try modifying the API comments in the code (e.g., in `cmd/api/main.go` or route handlers) to see how it affects the generated documentation. Then regenerate the docs:

```bash
make swag
# or
swag init -g main.go -d cmd/api,cmd/utils,cmd/models,cmd/db
```

## Running the Application

### Development Mode

To run the server with live reloading (if configured) or just standard run:

```bash
make dev
```

Or directly with Go:

```bash
go run ./cmd/api
```

The server will start at `http://localhost:5000` (or the port specified in `.env`).

### Docker Deployment

You can also run the application using Docker. This is useful for creating a consistent environment.

1.  **Build the Docker Image**:
    ```bash
    docker build -t go-gin-doc .
    ```

2.  **Run the Container**:
    Ensure your database is accessible from within the container. If you are running PostgreSQL on your host machine, you might need to use `host.docker.internal` instead of `localhost` in your `DB_URL`.

    ```bash
    # Example running on port 8080
    docker run -p 8080:8080 --env-file .env -e PORT=8080 go-gin-doc
    ```

    > **Note**: The Dockerfile exposes port 8080. We pass `-e PORT=8080` to ensure the application listens on the same port that Docker expects.

## Testing

Run the test suite:

```bash
make test
```

## Troubleshooting

-   **Database Connection Failed**: Double-check your `DB_URL` in `.env`. Ensure the PostgreSQL service is running and the database exists.
-   **Port Already in Use**: Change the `PORT` variable in `.env` to a free port (e.g., 5001).
-   **Migrate Command Not Found**: Ensure `golang-migrate` is installed and added to your system's PATH.
