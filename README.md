# Hexagonal API with Golang, Gin, and PostgreSQL

This project is a hexagonal architecture-based API built with Golang, utilizing the Gin web framework and PostgreSQL as the database. The architecture is designed to promote separation of concerns, making the codebase more maintainable and scalable.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

The project follows the hexagonal architecture (ports and adapters), ensuring a clear separation between the business logic and the outer layers like web frameworks and databases.

```plaintext
├── cmd/
│   └── api/
│       ├── bootstrap/
│       │   └── bootstrap.go
│       └── main.go
├── internal/
│   ├── platform/
│   │   ├── server/
│   │   │   ├── server.go
│   │   │   ├── handler/
│   │   │   │   └── health.go
│   │   ├── database/
│   │   │   ├── postgres.go       # Maneja la conexión a la base de datos y transacciones
│   │   └── repository/           # Implementaciones de los repositorios
│   │       └── postgres_user_repository.go   # Implementación del repositorio de usuarios en Postgres
│   ├── domain/
│   │   └── user/
│   │       ├── repository.go      # Interfaz del repositorio de usuario
│   │       └── service.go         # Lógica de negocio del usuario
│   └── app/
│       └── user/
│           └── controller.go      # Controlador que maneja las peticiones HTTP
├── go.mod
└── go.sum
```

## Getting Started

These instructions will help you set up the project on your local machine for development and testing purposes.

### Prerequisites

Ensure that you have the following installed on your machine:

- Golang (v1.18+)
- PostgreSQL
- Git

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Lalo64GG/Go-Hexagonal.git
   cd Go-Hexagonal

2. **Install dependecies:**
   ```bash
   go mod tidy

### Environment Variables

Create a `.env` file in the root directory of the project with the following content:

```bash
DB_CONN="user=yourusername password=yourpassword dbname=yourdbname sslmode=disable"

