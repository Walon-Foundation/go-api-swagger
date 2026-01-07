# Go Gin Doc

A robust API server built with Go and Gin for managing events and attendance. **This project is primarily a learning resource designed to explore how to integrate Swagger documentation with the Gin framework in Go.** It demonstrates a clean architecture approach to building RESTful APIs with Swagger documentation, JWT authentication, and PostgreSQL integration.

## 🚀 Features

- **User Authentication**: Secure signup and login endpoints using JWT.
- **Event Management**: Create and retrieve events.
- **Protected Routes**: Middleware to secure endpoints requiring authentication.
- **Swagger Documentation**: Interactive API documentation available at `/swagger/index.html`.
- **Database Integration**: PostgreSQL support with migration handling.

## 🎓 Learning Focus

This project serves as a practical playground for understanding:
- **Swagger/OpenAPI**: How to annotate Go code to automatically generate interactive API documentation.
- **Gin Framework**: Routing, middleware, and request handling in Gin.
- **Clean Architecture**: Structuring a Go application for maintainability and scalability.
- **Authentication**: Implementing JWT-based security flows.

## 🛠️ Tech Stack

- **Language**: [Go](https://go.dev/) (1.23+)
- **Framework**: [Gin Web Framework](https://gin-gonic.com/)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **Driver**: [pgx](https://github.com/jackc/pgx)
- **Authentication**: JWT (JSON Web Tokens)
- **Documentation**: [Swagger](https://swagger.io/) (swaggo)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)
- **Containerization**: [Docker](https://www.docker.com/)

## 📂 Project Structure

```
.
├── cmd
│   ├── api          # Main API application entry point and routes
│   ├── db           # Database connection and migrations
│   ├── models       # Data models and structures
│   └── utils        # Utility functions
├── docs             # Swagger documentation files
├── Makefile         # Build and run commands
├── Dockerfile       # Docker build configuration
├── .dockerignore    # Files excluded from Docker context
├── docker-compose.yml # Docker composition for dependencies
└── go.mod           # Go module definition
```

## 🏁 Getting Started

For detailed setup instructions, please refer to [setup.md](setup.md).

### Prerequisites

- Go 1.23 or higher
- PostgreSQL
- Make (optional, for running Makefile commands)

### Quick Start

1.  **Clone the repository**
    ```bash
    git clone https://github.com/Walon-Foundation/go-gin-doc.git
    cd go-gin-doc
    ```

2.  **Setup Environment Variables**
    Create a `.env` file in the root directory:
    ```env
    PORT=5000
    JWT_SECRET=your_super_secret_key
    DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
    ```

3.  **Run the Application**
    ```bash
    make dev
    ```
    Or manually:
    ```bash
    go run ./cmd/api
    ```

4.  **Access the API**
    - Server running at: `http://localhost:5000`
    - Swagger Docs: `http://localhost:5000/swagger/index.html`

5.  **Run with Docker (Optional)**
    ```bash
    # Build the image
    docker build -t go-gin-doc .

    # Run the container (ensure DB_URL points to a reachable database)
    docker run -p 8080:8080 --env-file .env -e PORT=8080 go-gin-doc
    ```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔒 Security

For security considerations and reporting vulnerabilities, please see [security.md](security.md).
