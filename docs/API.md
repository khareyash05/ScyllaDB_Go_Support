# API Documentation

REST API for managing trivia facts. Base URL: `/api`.

## Endpoints

### GET /api/health

Health check endpoint for load balancers and probes.

**Response:** `200 OK` — `{ "status": "ok", "service": "trivia-api", "version": "1.0.0", "timestamp": "2024-01-15T10:00:00Z" }`

---

### GET /api/facts/count

Returns the total number of facts. Useful for pagination UIs.

**Response:** `200 OK` — `{ "count": 42 }`

---

### GET /api/facts

List trivia facts with optional pagination.

| Parameter | In    | Type   | Required | Default | Description                          |
|-----------|-------|--------|----------|---------|--------------------------------------|
| limit     | query | integer | No       | 100     | Max facts to return (1–1000)         |
| offset    | query | integer | No       | 0       | Number of facts to skip (pagination) |
| sort      | query | string | No       | desc    | Sort order: `asc` (oldest first) or `desc` (newest first) |

**Response:** `200 OK` — JSON array of facts.

```json
[
  {
    "ID": 1,
    "question": "What is the capital of France?",
    "answer": "Paris",
    "createdAt": "2024-01-15T10:00:00Z",
    "updatedAt": "2024-01-15T10:00:00Z"
  }
]
```

---

### POST /api/facts

Create a new trivia fact.

**Request body:** JSON

| Field    | Type   | Required | Description         |
|----------|--------|----------|---------------------|
| question | string | Yes      | The trivia question |
| answer   | string | Yes      | The answer          |

**Response:** `201 Created` — The created fact (same schema as above).

**Errors:** `500 Internal Server Error` — Invalid JSON or server error.

---

### GET /api/facts/:id

Fetch a single fact by ID.

| Parameter | In   | Type | Required | Description      |
|-----------|------|------|----------|------------------|
| id        | path | uint | Yes      | Fact ID          |

**Response:** `200 OK` — The fact object.

**Errors:** `400 Bad Request` — Invalid ID. `404 Not Found` — Fact not found.

---

### PATCH /api/facts/:id

Update a fact by ID. Accepts partial updates (only provided fields are updated).

**Request body:** JSON (all fields optional)

| Field    | Type   | Required | Description         |
|----------|--------|----------|---------------------|
| question | string | No       | Updated question    |
| answer   | string | No       | Updated answer      |

**Response:** `200 OK` — The updated fact object.

**Errors:** `400 Bad Request` — Invalid ID or JSON. `404 Not Found` — Fact not found.

---

### DELETE /api/facts/:id

Delete a fact by ID.

| Parameter | In   | Type | Required | Description      |
|-----------|------|------|----------|------------------|
| id        | path | uint | Yes      | Fact ID          |

**Response:** `204 No Content` — Fact deleted.

**Errors:** `400 Bad Request` — Invalid ID. `404 Not Found` — Fact not found.

---

## OpenAPI Spec

The canonical API specification is in `openapi.yaml` at the project root. Use it for code generation, validation, or tooling.
