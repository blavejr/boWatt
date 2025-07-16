
# BoWatt API Docs
this is a quick overview of this tiny api, note there is a `postman_collection.json` file which you can import into postman and quickly get started `docker-compose up --build` will build the backed, frontend and mongodb (unseeded)

For ease I highly recommend importing the provided postman json file

Base URL: `http://localhost:8080`

---

### `GET /ping`

Check if server is alive.
**Response**: 
```json
{
    "message": "pong"
}
```

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

**Response**: 
```json
{
    "message": "User created",
    "token": "10ec1495-82fa-426f-8d16-34e3f8eb3b91"
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


**Response**: 
```json
{
    "message": "User created",
    "token": "10ec1495-82fa-426f-8d16-34e3f8eb3b91"
}
```

---

### `GET /profile`

Get user profile (requires auth).
**Headers**:

```
Authorization: <token>
```

```json
{
    "created_at": 1752497727,
    "username": "testuer"
}
```

---

### `GET /files`

List all uploaded files (requires auth).
**Headers**:

```
Authorization: <token>
```

```json
[
    {
        "ID": "6875122480a754ea7c7be6d9",
        "UserId": "68751200269c81db4402bfbb",
        "Name": "spidermonks.txt",
        "Content": "Sure! Here's a long, informative, and engaging blog post about **spider monkey...",
        "Timestamp": 1752502746
    }
]
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

```json
{
    "id": "daae7bb87d026c0bfb1cce764bdbd9b28bc4eb88b5f328daf93c77087b5922cc",
    "message": "You already uploaded this file"
}
```
---

### `POST /uploads`

Upload multiple files (requires auth).
**Headers**:

```
Authorization: <token>
```

**Form Data**:

```
files: <file>,<file>,<file>,<file>...
```

```json
{
    "failed": null,
    "message": "Upload completed",
    "skipped": null,
    "uploaded": [
        "file1.txt",
        "file2.txt",
        "file3.txt",
        "file4.txt"
        ...
    ]
}
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

```json
{
    "results": [
        "Renewable energy is transforming the way we power our world. Unlike fossil fuels, renewable sources such as solar, wind, hydro, and geothermal are naturally replenished and have a much lower environmental impact. By harnessing these resources, we can reduce greenhouse gas emissions, improve air quality, and create new jobs in the green economy.",
        "Solar panels and wind turbines are becoming more efficient and affordable, making clean energy accessible to more people than ever before. Governments and businesses are investing in renewable infrastructure, recognizing its potential to drive economic growth while protecting the planet.",
        "Transitioning to renewable energy is essential for combating climate change and ensuring a sustainable future for generations to come. By supporting clean energy initiatives, we can all play a part in building a healthier, more resilient world.",
        "The concept of harnessing natural forces for energy is not new. For centuries, humans have used windmills to grind grain and waterwheels to power machinery. However, the industrial revolution shifted the world’s focus to coal, oil, and gas, which fueled unprecedented growth but also led to environmental degradation. In recent decades, the urgent need to address climate change has reignited interest in renewable energy sources.",
        "Solar energy is one of the fastest-growing renewable sectors. Photovoltaic (PV) panels convert sunlight directly into electricity, while solar thermal systems use the sun’s heat for water heating and industrial processes. Advances in solar technology have dramatically reduced costs, making solar power competitive with traditional energy sources in many regions. Large-scale solar farms and rooftop installations are now common sights in both urban and rural landscapes.",
        "Wind Energy: Capturing Nature’s Kinetic Force",
        ...
    ]
}
```

---

### `GET /query/history`

Query the searches of a user from a file (requires auth).
**Headers**:

```
Authorization: <token>
```

```json
[
    {
        "ID": "687747587bf8a334178f1a88",
        "Query": "part",
        "FileHash": "2966ad1f5737fe9b93ae8b2ca1a4fcaea26033d75902376e64c0e85eb9cf9a07",
        "UserId": "6874fe3f1316ef1490e85d4f",
        "Timestamp": 1752647512
    }
]
```
