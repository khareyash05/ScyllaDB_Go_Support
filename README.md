# ScyllaDB Go Support

Trivia facts application with ScyllaDB/Go backend.

## Documentation

- **[API Reference](docs/API.md)** — Full REST API documentation
- **openapi.yaml** — OpenAPI 3.0 spec for codegen and tooling

## API Reference

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

```json
{
  "ID": 2,
  "question": "How many continents are there?",
  "answer": "7",
  "createdAt": "2024-01-15T10:05:00Z",
  "updatedAt": "2024-01-15T10:05:00Z"
}
```

---

### Web UI

- `GET /` - List facts (HTML)
- `GET /fact` - New fact form (HTML)
- `POST /fact` - Create fact (form submit, redirects to confirmation)
