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

### Validation
I did not get to adding validation to the routes but this would be a good improvement, validation on both the frontend input, backend routes and on the database writes, simply to sanitize data and make sure that everything is getting what it needs in the shape that it needs it in.

### Production environment
The above is for demo purposes only and the architecture in a full scale production environment would be very much different, first I would use ports/adapters architecture for a cleaner code base on the backend, this architecture not only seperates concerns well but also allows us to very easily swap out parts of the program such as databases, caches, etc

I would also seperate it into multiple services, auth service, file service and even a query service to further seperate concerns and decouple the project, that way if one part of it is down for any reason the rest stay up and are able to function.

The in-memory cache would be replaced with a more robust persistant cache like redis and I would use something like elastic search to have better full text searches on the database

for inter service communication I would use events and some kind of message broker like RabbitMQ or Kafka, this way services can go about their business and not have to worry about the others, once a service completes its job it simply sends an event to the broker notifying the next service where continue, how to get the data etc, files can be queued for hashing and even further processing to enhance the searchability

I would also include observability tools on both the back and frontend to log errors such as datadog or sentry and something like pagerduty to notify the team when things go wrong in production

for deployments I would suggest using a CI for testing, automated builds and deployments, something like github actions or gitlab CI if we need more control, CI would watch for commits to a master branch and run a pipeline thats checks out the latest changes, runs the tests and builds a the application, the container can be pushed to a container registry to make it easy to quickly roll back to previous versions that are tested and are able to run straight away
