
# Task Management System (Gin)

## Overview

This is a simple Task Management backend built using Go with the Gin web framework. It supports task creation, retrieval, updating, deletion, pagination, and filtering by status.

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
  "title": "Write README",
  "description": "Complete documentation",
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
```
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

Using docker compose
```bash
docker-compose up --build 
```

On local
```bash
go mod tidy
go run cmd/main.go
```

## Microservices Concepts Demonstrated

- **Single Responsibility Principle**: Separate handler, service, and repository layers
- **Scalable architecture**: containerized and deployable using orchestration tools and can be horizontally scaled using replicas
- **Extensibility**: Ready for integration with other microservices (e.g., user/auth service)
