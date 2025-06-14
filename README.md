# Task Manager API in Go
A simple and lightweight Task Manager API written in Go. This project uses an in-memory database and is structured using a feature-based architecture. It is ideal for learning purposes and rapid prototyping of RESTful services in Go.

## Features
- Task management with basic CRUD operations (extendable)
- Feature-based project structure for scalability
- User authentication using JWT
- HTTP server built with Go's standard `net/http` package
- In-memory data storage

## Project Structure
```bash
task-manager-api/
├── cmd/                         # Entry points for different executables
│   └── server/                  # HTTP server 
│       └── main.go              # Application bootstrap: starts the server and loads routes
├── internal/                    # Internal application code (encapsulated, not imported elsewhere)
│   ├── task/                    # Task feature (CRUD operations, business logic, routing)
│   │   ├── handler.go           # HTTP handlers for task endpoints
│   │   ├── service.go           # Business logic and use-case layer for tasks
│   │   ├── repository.go        # In-memory data store logic for tasks
│   │   ├── model.go             # Data models (structs) for tasks
│   │   └── routes.go            # Route definitions specific to task endpoints
│   ├── user/                    # User feature (auth, registration, login)
│   │   ├── handler.go           # HTTP handlers for user endpoints
│   │   ├── service.go           # Business logic and use-case layer for users
│   │   ├── repository.go        # In-memory data store logic for users
│   │   ├── model.go             # Data models (structs) for users
│   │   └── routes.go            # Route definitions specific to user endpoints
├── utils/                       # Utility/helper functions used across the app
│   └── utils.go                 # Common helper methods (e.g., JWT generation, validation)
├── go.mod    
```

## API Endpoints

| Method | Endpoint        | Description                           | Auth Required |
|--------|-----------------|---------------------------------------|---------------|
| POST   | `/login`        | login user (generate jwt token)       | ✅            |
| POST   | `/register`     | register user                         | ✅            |
| GET    | `/logout`       | Logout user                           | ✅            |
| GET    | `/tasks`        | Retrieve user tasks                   | ✅            |
| POST   | `/tasks`        | add task                              | ✅            |
| GET    | `/task/{id}`    | Retrieve user task                    | ✅            |
| PATCH  | `/task/{id}`    | Update user task                      | ✅            |

> More endpoints can be added as needed.

### Prerequisites
- Go 1.22+ installed
- Git

### Run the API
1. Clone the repository:
   ```bash
   git clone https://github.com/harranali/task-manager-api.git
   cd task-manager-api
   ```
2. Running the API
   ```
   go run cmd/server/main.go
   ```

### License
This project is open-source under the MIT License.
