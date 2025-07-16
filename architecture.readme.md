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
  - Auth middleware checks for a valid token in `Authorization` header, I use a simple token but I could use JWT tokens that expires and a second refresh token in the http-only cookie, the refresh token will be much longer living the the Bearer token and will be used to reauthenticate and generate a new token for the user
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

## Deployments
Though I only imagine this app will only ever live on local computers I can think up how it could be deployed.
Ideally I would go with a CI/CD approach, this would require that tests are written to enable automatic build and deployment with each commit/merge to a master branch. something like gitlab CI or github actions could be used here, the pipeline would checkout the code, build it and run the tests, if all passes, we could push the container into a registry before checking it out and deploying it onto the available infra.
