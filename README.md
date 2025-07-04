# Go Students API

The **Go Students API** is a simple, high-performance RESTful API built using the Go programming language. It provides basic student management functionality including creating, retrieving, updating, and deleting student records.

## Features

- ✅ Add new student records
- ✅ Retrieve single or multiple student records
- ✅ Update existing student information
- ✅ Delete student records
- ✅ Fast and lightweight using Go

## Technologies Used

- Go (Golang)
- Standard Library (`net/http`, `encoding/json`, etc.)
- Optional: Any SQL/NoSQL database (sQlite)

## API Endpoints

| Method | Endpoint          | Description               |
|--------|-------------------|---------------------------|
| GET    | `/api/students`       | Get all students          |
| GET    | `/api/students/{id}`  | Get student by ID         |
| POST   | `/api/students`       | Create new student        |
| PUT    | `/api/students/{id}`  | Update student by ID      |
| DELETE | `/api/students/{id}`  | Delete student by ID      |

## Example Student JSON

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "johndoe@example.com",
  "age": 20
}
