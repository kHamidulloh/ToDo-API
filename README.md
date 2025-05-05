# ToDo REST API

This is a simple ToDo REST API built with Golang using the Gin web framework. It supports user registration, authentication with JWT, and CRUD operations for tasks. Task statuses are validated and must be either `pending` or `completed`.

## Features

- User registration and login with password hashing (bcrypt)
- JWT-based authentication and middleware
- Create, read, update, delete tasks
- Filter tasks by status (`pending` or `completed`)
- PostgreSQL database integration using GORM

## Technologies Used

- Go
- Gin
- GORM
- PostgreSQL
- JWT (github.com/dgrijalva/jwt-go)
- bcrypt
- godotenv

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/todo-api.git
cd todo-api

2. **Install dependencies**

go mod tidy

3. **Create a .env file**

Create a file named .env in the root directory and add the following:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo

4. **Run PostgreSQL locally**

Make sure PostgreSQL is running and the todo database is created:

CREATE DATABASE todo;

5. **Run the project**

go run ./src/main.go

API Endpoints
Auth
Register a new user

POST /register
Content-Type: application/json

{
  "username": "admin",
  "password": "password123"
}

Login

POST /login
Content-Type: application/json

{
  "username": "admin",
  "password": "password123"
}

Returns a JWT token:

{
  "token": "your-jwt-token"
}

    Use this token in the Authorization header for the following endpoints.

Tasks

All endpoints below require a valid JWT token in the Authorization header.

Example:

Authorization: your-jwt-token

Create a task

POST /tasks
Content-Type: application/json

{
  "title": "Buy milk",
  "description": "Buy 2 liters of milk",
  "status": "pending"
}

Get tasks

GET /tasks

You can filter by status:

GET /tasks?status=completed

Update a task

PUT /tasks/:id
Content-Type: application/json

{
  "title": "Buy bread",
  "description": "Whole grain",
  "status": "completed"
}

Delete a task

DELETE /tasks/:id

Folder Structure

todo-api/
│
├── src/
│   ├── main.go             # Entry point
│   ├── models/             # GORM models (User, Task)
│   ├── handlers/           # Route handlers
│   ├── services/           # JWT and helper services
│   └── database/           # DB connection setup
│
├── go.mod
├── .env
└── README.md

License

This project is open source and free to use.
