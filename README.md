# ScyllaDB Go Support

Trivia facts application with ScyllaDB/Go backend.

## Documentation

- **[API Reference](docs/API.md)** — Full REST API documentation
- **openapi.yaml** — OpenAPI 3.0 spec for codegen and tooling

## API Reference

### GET /api/health

Health check. Returns `{ "status": "ok", "service": "trivia-api", "version": "1.0.0" }`.

---

### GET /api/facts/count

Returns total fact count. `{ "count": N }`

---

### GET /api/facts

List all trivia facts.

**Parameters:**

| Name   | In    | Type    | Required | Description                                      |
|--------|-------|---------|----------|--------------------------------------------------|
| limit  | query | integer | No       | Max facts to return (default 100, max 1000)      |
| offset | query | integer | No       | Number of facts to skip for pagination (default 0)|

**Output:** JSON array of facts.

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

**Parameters:** (JSON body)

| Name     | Type   | Required | Description        |
|----------|--------|----------|--------------------|
| question | string | Yes      | The trivia question|
| answer   | string | Yes      | The answer         |

**Output:** The created fact (201 Created) or error (500).

---

### GET /api/facts/:id

Fetch a single fact by ID. Returns 200 with the fact, or 400/404 on error.

---

### PATCH /api/facts/:id

Update a fact by ID. Accepts partial JSON body (question, answer). Returns 200 with updated fact.

---

### DELETE /api/facts/:id

Delete a fact by ID. Returns 204 No Content on success, or 400/404 on error.

---
    </content>
  </file>
  <file path="docs/API.md">
    <content>
# API Documentation

REST API for managing trivia facts. Base URL: `/api`.

## Endpoints

### GET /api/health

Health check endpoint for load balancers and probes.

**Response:** `200 OK` — `{ "status": "ok", "service": "trivia-api", "version": "1.0.0" }`

---

### GET /api/facts/count

Returns the total number of facts. Useful for pagination UIs.

**Response:** `200 OK` — `{ "count": 42 }`

---

### GET /api/facts

List trivia facts with optional pagination.

| Parameter | In    | Type    | Required | Default | Description                          |
|-----------|-------|---------|----------|---------|--------------------------------------|
| limit     | query | integer | No       | 100     | Max facts to return (1–1000)         |
| offset    | query | integer | No       | 0       | Number of facts to skip (pagination) |

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
    </content>
  </file>
  <file path="openapi.yaml">
    <content>
openapi: 3.0.3
info:
  title: ScyllaDB Go Support API
  description: API for managing trivia facts.
  version: 1.0.0

paths:
  /api/health:
    get:
      summary: Health check
      description: Returns API health status for load balancers and probes.
      responses:
        "200":
          description: Service is healthy
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: string
                    example: ok
                  service:
                    type: string
                    example: trivia-api
                  version:
                    type: string
                    example: "1.0.0"
                    description: API version
  /api/facts/count:
    get:
      summary: Count facts
      description: Returns the total number of facts. Useful for pagination.
      responses:
        "200":
          description: Total count
          content:
            application/json:
              schema:
                type: object
                properties:
                  count:
                    type: integer
                    example: 42
  /api/facts:
    get:
      summary: List facts
      description: List trivia facts with optional pagination.
      parameters:
        - name: limit
          in: query
          description: Max facts to return
          schema:
            type: integer
            default: 100
            minimum: 1
            maximum: 1000
        - name: offset
          in: query
          description: Number of facts to skip for pagination (default 0)
          schema:
            type: integer
            default: 0
            minimum: 0
      responses:
        "200":
          description: List of facts
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Fact"
    post:
      summary: Create a fact
      description: Creates a new trivia fact.
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/FactInput"
      responses:
        "201":
          description: Fact created
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Fact"
        "500":
          description: Server error
  /api/facts/{id}:
    get:
      summary: Get a fact by ID
      description: Returns a single trivia fact by its ID.
      parameters:
        - name: id
          in: path
          required: true
          description: Fact ID
          schema:
            type: integer
            minimum: 1
      responses:
        "200":
          description: The fact
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Fact"
        "400":
          description: Invalid ID
        "404":
          description: Fact not found
    patch:
      summary: Update a fact
      description: Updates a trivia fact by ID. Accepts partial updates.
      parameters:
        - name: id
          in: path
          required: true
          description: Fact ID
          schema:
            type: integer
            minimum: 1
      requestBody:
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/FactInput"
      responses:
        "200":
          description: The updated fact
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Fact"
        "400":
          description: Invalid ID or JSON
        "404":
          description: Fact not found
    delete:
      summary: Delete a fact
      description: Deletes a trivia fact by ID.
      parameters:
        - name: id
          in: path
          required: true
          description: Fact ID
          schema:
            type: integer
            minimum: 1
      responses:
        "204":
          description: Fact deleted
        "400":
          description: Invalid ID
        "404":
          description: Fact not found
components:
  schemas:
    Fact:
      type: object
      properties:
        ID:
          type: integer
        question:
          type: string
        answer:
          type: string
        createdAt:
          type: string
          format: date-time
        updatedAt:
          type: string
          format: date-time
    FactInput:
      type: object
      required:
        - question
        - answer
      properties:
        question:
          type: string
        answer:
          type: string
    