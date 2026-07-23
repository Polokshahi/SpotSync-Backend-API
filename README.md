# SpotSync Backend API

SpotSync is a RESTful backend API built with **Go**, **Echo Framework**, **PostgreSQL (Neon)**, **GORM**, and **JWT Authentication**. It provides user authentication, zone management, and reservation management using a clean layered architecture.



# Features

* User Registration
* User Login
* JWT Authentication
* Password Hashing with bcrypt
* Protected Routes using Middleware
* Zone Management
* Reservation Management
* PostgreSQL Database Integration
* GORM ORM
* Layered Architecture (Handler → Service → Repository)

---

# Tech Stack

* Go
* Echo Framework
* PostgreSQL (Neon)
* GORM
* JWT
* bcrypt
* godotenv

---

# Project Structure

```text
spotsync/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   ├── dto/
│   ├── handler/
│   ├── models/
│   ├── repository/
│   └── service/
├── pkg/
│   └── middleware/
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

# API Endpoints

## Authentication

| Method | Endpoint                | Description         |
| ------ | ----------------------- | ------------------- |
| POST   | `/api/v1/auth/register` | Register a new user |
| POST   | `/api/v1/auth/login`    | Login user          |

---

## Zones

| Method | Endpoint            | Description       |
| ------ | ------------------- | ----------------- |
| GET    | `/api/v1/zones`     | Get all zones     |
| GET    | `/api/v1/zones/:id` | Get zone by ID    |
| POST   | `/api/v1/zones`     | Create a new zone |

---

## Reservations

| Method | Endpoint                               | Description                       |
| ------ | -------------------------------------- | --------------------------------- |
| POST   | `/api/v1/reservations`                 | Create reservation                |
| GET    | `/api/v1/reservations`                 | Get all reservations              |
| GET    | `/api/v1/reservations/my-reservations` | Get logged-in user's reservations |
| DELETE | `/api/v1/reservations/:id`             | Delete reservation                |

---

# Authentication

Protected endpoints require a JWT token.

Example Header:

```http
Authorization: Bearer YOUR_JWT_TOKEN
```

---

# Installation

## Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/spotsync.git
```

```bash
cd spotsync
```

---

## Install Dependencies

```bash
go mod tidy
```

---

## Create Environment File

Create a `.env` file in the project root.

Example:

```env
DATABASE_URL=your_neon_database_url
JWT_SECRET=your_secret_key
PORT=8080
```

---

## Run the Application

```bash
go run ./cmd/api
```

If everything is configured correctly, the server will start successfully.

---

# Build

```bash
go build ./...
```

---

# Code Quality

```bash
go fmt ./...
go vet ./...
```

---

# Sample Register Request

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456",
  "role": "user"
}
```

---

# Sample Login Request

```json
{
  "email": "john@example.com",
  "password": "123456"
}
```

---

# Sample Login Response

```json
{
  "token": "YOUR_JWT_TOKEN"
}
```

---

# Architecture

The project follows a layered architecture.

```
Client
   │
HTTP Request
   │
Handler
   │
Service
   │
Repository
   │
PostgreSQL (Neon)
```

### Handler

Handles HTTP requests and responses.

### Service

Contains business logic.

### Repository

Communicates with the database using GORM.

### Models

Represents database tables.

### DTO

Defines request and response payloads.

### Middleware

Protects routes using JWT authentication.

---

# Security

* Passwords are hashed using bcrypt.
* JWT is used for authentication.
* Environment variables are stored outside the source code.
* Protected routes require a valid JWT token.

---

# Dependencies

* Echo Framework
* GORM
* PostgreSQL Driver
* JWT
* bcrypt
* godotenv

---

# Author

**MD. Abdulla Hel Shahi (Polok Shahi)**

GitHub: https://github.com/Polokshahi

LinkedIn: https://www.linkedin.com/in/polokshahi/

---

# License

This project was developed as an academic assignment for learning backend development using Go.
