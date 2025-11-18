# Todo Task API

A simple RESTful API for managing todo items built with Go and Gorilla Mux.

## Features

- Create, read, update, and delete todos
- JSON API responses
- RESTful architecture
- In-memory data storage

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. Clone the repository
```bash
git clone https://github.com/azhari-zaidie/golang-todo-task.git
cd golang-todo-task
```

2. Install dependencies
```bash
go mod tidy
```

3. Run the application
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/todos` | Get all todos |
| GET | `/api/todos/{id}` | Get a specific todo |
| POST | `/api/todos` | Create a new todo |
| PUT | `/api/todos/{id}` | Update a todo |
| DELETE | `/api/todos/{id}` | Delete a todo |
| GET | `/health` | Health check |

## API Testing

Test the API using Postman:

[View Postman Collection](https://www.postman.com/docking-module-candidate-18013215/workspace/learning/collection/40777438-207cd416-324c-49a7-b4b7-2710daf5220b?action=share&creator=40777438&active-environment=40777438-f84441fd-6beb-4b71-b3bc-af1665d7e17c)

### Example Requests

**Create a Todo**
```json
POST /api/todos
{
  "title": "Learn Go",
  "description": "Master Go programming language",
  "completed": false
}
```

**Update a Todo**
```json
PUT /api/todos/1
{
  "title": "Learn Go",
  "description": "Master Go programming language",
  "completed": true
}
```

## Project Structure

```
.
├── handlers/          # HTTP request handlers
├── models/           # Data models
├── routes/           # Route definitions
├── main.go           # Application entry point
├── go.mod            # Go module dependencies
└── README.md         # Project documentation
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gorilla Mux](https://github.com/gorilla/mux) - HTTP router and URL matcher

## License

This project is open source and available under the MIT License.
