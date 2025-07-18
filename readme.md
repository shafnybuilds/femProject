# Workout Tracker - Go Backend API

A RESTful API backend built with Go, using Chi router and PostgreSQL database for workout management.

## 🏗️ Project Structure

```
workout-tracker-go/
├── main.go                     # Application entry point
├── docker-compose.yml          # PostgreSQL database setup
├── go.mod                      # Go module dependencies
├── internal/
│   ├── app/
│   │   └── app.go             # Application configuration & health check
│   ├── api/
│   │   └── workout_handler.go  # Workout HTTP handlers
│   └── routes/
│       └── routes.go          # Route definitions
└── database/
    └── postgres-data/         # PostgreSQL data volume
```

## 🚀 Features

- **RESTful API** for workout management
- **PostgreSQL** database integration
- **Chi router** for HTTP routing
- **Configurable port** via command-line flags
- **Health check** endpoint
- **Docker Compose** for easy database setup

## 📋 API Endpoints

| Method | Endpoint         | Description           |
| ------ | ---------------- | --------------------- |
| `GET`  | `/health`        | Health check endpoint |
| `GET`  | `/workouts/{id}` | Get workout by ID     |
| `POST` | `/workouts`      | Create new workout    |

## 🛠️ Prerequisites

- **Go 1.23.5+**
- **Docker & Docker Compose**
- **Git**

## 🚀 Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/shafnybuilds/femProject.git
cd femProject
```

### 2. Start the database

```bash
docker compose up -d
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the application

```bash
# Default port 8080
go run main.go

# Custom port
go run main.go -port=3000
```

## 🧪 Testing the API

### Health Check

```bash
curl http://localhost:8080/health
```

### Get Workout by ID

```bash
curl http://localhost:8080/workouts/123
```

### Create Workout

```bash
curl -X POST http://localhost:8080/workouts
```

## 🐳 Docker Services

    - Demo Configs -

- **PostgreSQL 12.4-alpine** - Database server
- **Port**: \*\*\*\*
- **Database**: \*\*\*\*
- **User**: \*\*\*\*
- **Password**: \*\*\*\*

## 🔧 Configuration

### Environment Variables

```bash
# Database connection (when implemented)
DB_HOST=****
DB_PORT=****
DB_USER=****
DB_PASSWORD=****
DB_NAME=****
```

### Command Line Flags

```bash
go run main.go -port=8080  # Set custom port
```

## 📦 Dependencies

- **[Chi Router](https://github.com/go-chi/chi/v5)** - HTTP router and URL matcher
- **PostgreSQL** - Database (via Docker)

## 🏃‍♂️ Development

### Project follows Go best practices:

- **Clean architecture** with separated concerns
- **Handler pattern** for HTTP endpoints
- **Dependency injection** through Application struct
- **Error handling** with proper HTTP status codes

## 📝 License

This project is licensed under the MIT License.

---

**Author**: [@shafnybuilds](https://github.com/shafnybuilds)
