
# Task Management System (Gin)

## Overview

This is a simple Task Management backend built using Go with the Gin web framework. It supports task creation, retrieval, updating, deletion, pagination, and filtering by status.

## Prerequisites

Before you run this project, ensure you have the following installed:

- [Go 1.20+] (if running via local)
- [Docker]
- [Docker Compose] (if running via `docker-compose`)
- Git

## 🔧 Problem Breakdown & Design Decisions

### Requirements:
- RESTful API for basic CRUD operations on tasks.
- Pagination support for listing tasks.
- Filtering support by task status.
- Modular structure to demonstrate microservices architecture principles.
- Containerized using Docker and orchestrated using docker-compose

### Design Decisions:
- **Framework**: Gin was chosen for its performance and simplicity in building REST APIs.
- **ORM**: GORM was used for easy DB integration and model handling.
- **Database**: SQLite for simplicity in local development.
- **Architecture**: Layered with clear separation of concerns:
  - `handlers`: Handle HTTP requests/responses.
  - `services`: Contain core business logic.
  - `repository`: Interact with the database layer.
  - `models`: Define data entities.
- **Containerization** : Containerized using docker and orchestrated via docker-compose , can be extended to kubernetes and AWS ECS for **fault tolerance** and **scalability**  
- **Extensibility**: Designed to integrate smoothly with user/auth microservices later.

## API Endpoints

| Method | Endpoint       | Description            |
|--------|----------------|------------------------|
| POST   | /tasks         | Create a new task      |
| GET    | /tasks         | List tasks             |
| GET    | /tasks/:id     | Get a task by ID       |
| PUT    | /tasks/:id     | Update a task          |
| DELETE | /tasks/:id     | Delete a task          |


## Sample Request and responses

### POST /tasks

**Description:** Create a new task

**Request Body (`application/json`):**

```json
{
  "title": "Finish project",
  "description": "Complete the task management service",
  "status": "Pending"
}
```

**Response :**
```json
HTTP status : 201 Created

{
  "id": 1,
  "title": "Finish project",
  "description": "Complete the task management service",
  "status": "Pending",
  "created_at": "2025-05-04T12:00:00Z",
  "updated_at": "2025-05-04T12:00:00Z"
}
```

### GET /tasks — List tasks (with optional filtering & pagination)

#### Query Parameters (GET /tasks)

- `status` — filter by task status
- `page` — page number (default: 1)
- `limit` — number of items per page (default: 10)

Get **All Tasks**
```
GET /tasks 
```

Get tasks by **Task Id**
```
GET /tasks/1 
```

Get tasks using **filter and pagination**
```
GET /tasks?status=Pending&page=1&limit=2
```

**Response**
```json
Http status : 200 OK
[
  {
    "id": 1,
    "title": "Finish project",
    "description": "Complete the task management service",
    "status": "Pending",
    "created_at": "2025-05-04T12:00:00Z",
    "updated_at": "2025-05-04T12:00:00Z"
  },
  {
    "id": 2,
    "title": "Write README",
    "description": "Add detailed documentation",
    "status": "Pending",
    "created_at": "2025-05-04T12:30:00Z",
    "updated_at": "2025-05-04T12:30:00Z"
  }
]

```

### PUT /tasks/:id 
**Description:** Update a task

**Request Body**

```json
PUT /tasks/1
Content-Type: application/json

{
  "title": "Finish backend project",
  "description": "Complete with Docker and README",
  "status": "Completed"
}
```
**Response**

```json
Http status : 200 OK
{
  "id": 1,
  "title": "Finish backend project",
  "description": "Complete with Docker and README",
  "status": "Completed",
  "created_at": "2025-05-04T12:00:00Z",
  "updated_at": "2025-05-04T13:00:00Z"
}
```

### DELETE /tasks/:id 

**Description:** Delete a task

**Request**
```bash
DELETE /tasks/1
```

**Response**
```bash
Http status : 204 StatusNoContent
```

## Running the Service

The service will be available on `localhost:8080`

Using docker compose
```bash
docker-compose up --build 
```

On local
```bash
go mod tidy
go run cmd/main.go
```

### Improvements
- Currently, the same `Task` model is used for both request parsing and response formatting.
- It is a better practice to use **DTOs (Data Transfer Objects)** for handling incoming requests and **DAOs (Data Access Objects)** or models for database operations. Exposing internal models directly can lead to tight coupling and security issues.
- Error handling can be improved — currently, many endpoints return generic errors. Structured and descriptive error responses would improve debuggability and client usability.
- Future improvements may include:
  - Adding middleware for structured logging and request tracing
  - Integrating with an external authentication/authorization service
  - Switching from SQLite to PostgreSQL or another production-ready database
  - Adding rate limiting and request validation
  - Including Swagger/OpenAPI documentation for automated API exploration



## Microservices Concepts Demonstrated

- **Single Responsibility Principle**: Clean separation between handlers (controllers), services (business logic), and repositories (data access).
- **Scalable Architecture**: The service is containerized using Docker and can be deployed via orchestration tools like Docker Compose/Swarm or Kubernetes. It supports horizontal scaling via replicas.
- **Extensibility**: Designed to support the addition of new microservices (e.g., User Service, Auth Service). Inter-service communication could be enabled via REST APIs or gRPC in future iterations.
- **Environment Isolation**: Can Use environment variables and a `.env`-friendly setup to allow flexible deployment across environments (dev/staging/production).
- **Stateless Design**: Each replica of the service can handle any request independently, making it suitable for load balancing and distributed deployments.
