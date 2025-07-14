
# BoWatt API Docs
this is a quick overview of this tiny api, note there is a `postman_collection.json` file which you can import into postman and quickly get started `docker-compose up --build` will build the backed, frontend and mongodb (unseeded)

Base URL: `http://localhost:8080`

---

### `GET /ping`

Check if server is alive.
**Response**: `pong`

---

### `POST /signup`

Register a new user.
**Body (JSON)**:

```json
{
  "username": "string",
  "password": "string"
}
```

---

### `POST /login`

Login to get an auth token.
**Body (JSON)**:

```json
{
  "username": "string",
  "password": "string"
}
```

---

### `GET /profile`

Get user profile (requires auth).
**Headers**:

```
Authorization: <token>
```

---

### `GET /files`

List all uploaded files (requires auth).
**Headers**:

```
Authorization: <token>
```

---

### `POST /upload`

Upload a file (requires auth).
**Headers**:

```
Authorization: <token>
```

**Form Data**:

```
file: <file>
```

---

### `POST /query`

Query content from a file (requires auth).
**Headers**:

```
Authorization: <token>
```

**Body (JSON)**:

```json
{
  "FileHash": "string",
  "query": "string"
}
```

---

Let me know if you want to add example responses or status codes.
