# Go Backend Development – User Management API

A RESTful backend service built using Go, Fiber, PostgreSQL, SQLC, and Uber Zap.
This project manages users with name and date of birth, calculates age dynamically,
and supports pagination and structured logging.

Before you go further , you can check this API directly without any setup at  : http://54.225.177.15 

---

## Tech Stack

- Go
- Fiber
- PostgreSQL
- SQLC
- Uber Zap

---

## Features

- Create, Read, Update, Delete users
- Dynamic age calculation (age is not stored in DB)
- Pagination for listing users
- Request ID injection in responses
- Request duration logging
- Structured logging using Uber Zap

---

## Prerequisites

- Go (version 1.20 or higher)
- PostgreSQL running locally
- sqlc installed

---

## Database Setup (If not using Docker)

Login into your PostgreSQL:
```bash
psql -U admin
```


Create database and table:

```sql
CREATE DATABASE user_dob_db;

\c user_dob_db;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

```

---

## Project Setup

1. Clone the repository
2. Create `.env` file in git directory.
```bash
DATABASE_URL=postgres://YOURUSERNAME:YOURPASSWORD@postgres:5432/user_dob_db?sslmode=disable
SERVER_PORT=3000
```
Change `postgres:5432` (For Docker) to `localhost:5432` if you are going to use local setup below. 
Also change `YOURUSERNAME` to your username & password respectively.

3. Generate SQLC code:

Here is the `.env` file conten:


```bash
sqlc generate
```

---

## Running the Application (Locally)

From the project root, first download dependencies using:

```bash
go mod download
```

```bash
go run cmd/server/main.go
```

The server will start on:

```
http://localhost:3000
```
---

## Running the Application with Docker

Just go to project directory and hit the following command:

```bash
docker compose build --no-cache
```
After completion of build step, run this command:

```bash
docker compose up -d
```

The server will start on:

```
http://localhost:3000
```

`Note: Docker compose will take care of initial DB table setup. You don't need to make table or db. Everything will be taken care of internally.`

---

## API Endpoints

| Method | Endpoint | Description |
|------|---------|-------------|
| POST | /users | Create a new user |
| GET | /users/:id | Get user by ID (returns age) |
| PUT | /users/:id | Update user |
| DELETE | /users/:id | Delete user |
| GET | /users?page=&limit= | List users with pagination |

POST : `/users` : Request Body:
```bash
{
    "name" : "Satyam",
    "dob" : "2001-04-05"
}
```
Same body can be used for PUT request of `/users/:id` endpoint.

---

## Testing Age Service Layer

Tests are written inside `./internal/service/age_test.go`.
Just hit to test age calculation logic:
```bash
go test ./internal/service
```
---