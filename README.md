---

## 🧩 Golang REST API with JWT Auth & SQLite

This project is a fully functional Event-Registration REST API built with **Go (Golang)** using the **Gin framework**, **JWT-based authentication**, and **SQLite** as the database. It provides secure user authentication and full CRUD functionality for managing events.

---

### ✨ Features

* 🔐 **JWT Authentication** (`/signup`, `/login`)
* 🧾 **Event CRUD API** (Create, Read, Update, Delete)
* ✅ **Protected Routes** using middleware
* 💾 **SQLite** database integration
* 🛡️ Secure password hashing with **bcrypt**
* ⚙️ Clean modular structure with packages

---

### 📁 Project Structure

```
├── db/              → Database setup & table creation
├── models/          → User & Event models with DB logic
├── routes/          → API route handlers
├── middlewares/     → JWT authentication middleware
├── utils/           → Token generation & password hashing
├── main.go          → App entry point
```

---

### 📮 API Endpoints

#### Public

* `POST /signup` – Create user account
* `POST /login` – Authenticate user and return JWT
* `GET /events` – List all events
* `GET /events/:id` – Get a single event

#### Protected (Requires Token)

* `POST /events` – Create new event
* `PUT /events/:id` – Update an event
* `DELETE /events/:id` – Delete an event

---

### 🛠️ Tech Stack

* **Golang**
* **Gin Gonic** (HTTP Web Framework)
* **SQLite3** (lightweight DB)
* **JWT** (Authentication)
* **bcrypt** (Password hashing)

---

### 🚀 Getting Started

```bash
go mod tidy
go run .
```

Server will start on: `http://localhost:8080`

---
