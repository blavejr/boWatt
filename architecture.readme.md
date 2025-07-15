# Backend Architecture Overview

Here’s a high-level overview of its structure and responsibilities:

---

## 1. Web Framework & Entry Point

- **Framework:** Uses [Gin](https://github.com/gin-gonic/gin) for HTTP routing and middleware.
- **Entry Point:** `main.go`
  - Sets up logging
  - Connects to MongoDB
  - Initializes controllers
  - Configures routes

- **Endpoints:**
  - **Public:** `signup`, `login`
  - **Protected:** `file upload`, `query`, `profile`

- **Middleware:**
  - Applies CORS for frontend-backend communication

---

## 2. Controllers (Business Logic Layer)

- **Location:** `controllers/`
- **Files:**
  - `users.go`: User registration, login, password hashing, token management
  - `files.go`: File uploads, deduplication (by hash), list files per user
  - `queries.go`: Search queries on uploaded files, including caching and fuzzy search
  - `helpers.go`: Utility for extracting the current user from context

---

## 3. Middleware

- **Location:** `middleware/auth.go`
- **Responsibilities:**
  - Auth middleware checks for a valid token in `Authorization` header
  - Loads user from DB and attaches to request context for use in controllers

---

## 4. Models (Data Layer)

- **Location:** `models/`
- **Files:**
  - `db.go`: MongoDB connection and collection helpers
  - `users.go`: Defines the `User` schema
  - `files.go`: Defines the `FileMetadata` schema

---

## 5. Utilities

- **Location:** `utils/`
- **Files:**
  - `cache.go`: In-memory cache for query results
  - `crypto.go`: File hashing (SHA-256)
  - `search.go`: Fuzzy search for file content
  - `storage.go`: In-memory file storage helpers

---

## 6. Persistence

- **Database:** MongoDB (connection managed in `models/db.go`)
- **Collections:**
  - `users`
  - `files`
  - `queries` (for search results and history)

---

## 7. Deployment

- **Dockerized:**
  - Includes `Dockerfile` and `Dockerfile.dev` for production and development
- **Entrypoint:**
  - Exposes port `8080` for the API server

---

## Request Flow Example

### User Signup/Login:
- Handled via public endpoints (`users.go`)
- Passwords hashed and tokens generated for auth

### Authenticated Requests:
- Protected by `AuthMiddleware`
- Loads user context for downstream handlers

### File Upload/Search:
- Handled via `files.go` and `queries.go`
- Files hashed, deduplicated, and stored in MongoDB
- Search uses fuzzy matching and results are cached
