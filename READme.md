# 🎵 MusicList-API

A RESTful Music Playlist API built in **Go** using the **Echo framework** with **PostgreSQL** as the database. It supports user authentication using **JWT**, structured logging via **Logrus**, and follows a modular architecture.

---

## 🚀 Features

- User Registration & Login
- JWT-Based Authentication
- CRUD operations on Songs (Create, Read, Update, Delete)
- PostgreSQL as the relational database
- Structured logs using Logrus
- Clean modular structure (config, routes, controllers, middleware, etc.)

---

## 🛠 Tech Stack & Libraries

| Purpose            | Package                                      |
|--------------------|----------------------------------------------|
| Web Framework      | [`github.com/labstack/echo/v4`](https://echo.labstack.com/) |
| JWT Auth           | [`github.com/golang-jwt/jwt/v4`](https://github.com/golang-jwt/jwt) |
| PostgreSQL Driver  | [`github.com/lib/pq`](https://pkg.go.dev/github.com/lib/pq) |
| DB Management      | `database/sql` from Go standard library      |
| Environment Loader | [`github.com/joho/godotenv`](https://github.com/joho/godotenv) |
| Logger             | [`github.com/sirupsen/logrus`](https://github.com/sirupsen/logrus) |

---

## 🧾 API Endpoints

### Public Routes

| Method | Endpoint     | Description          |
|--------|--------------|----------------------|
| GET    | `/`          | Welcome route        |
| POST   | `/register`  | Register a new user  |
| POST   | `/login`     | Login and get JWT    |

### Protected Routes (JWT Required)

| Method | Endpoint        | Description          |
|--------|------------------|----------------------|
| POST   | `/songs`         | Add a new song       |
| GET    | `/songs`         | List all songs       |
| PUT    | `/songs/:id`     | Update a song        |
| DELETE | `/songs/:id`     | Delete a song        |

🛡️ Protected routes require a JWT token in the `Authorization` header: