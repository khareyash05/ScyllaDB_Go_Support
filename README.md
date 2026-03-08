# ScyllaDB Go Support

Trivia facts application with ScyllaDB/Go backend.

## Features

- RESTful API to list and create trivia facts
- Pagination support via `limit` and `offset` query parameters
- Built-in Web UI for browser-based interactions
- High-performance ScyllaDB database integration

## API Reference

### GET /api/facts

List all trivia facts.

**Parameters:**

| Name   | In    | Type    | Required | Description                                      |
|--------|-------|---------|----------|--------------------------------------------------|
| limit  | query | integer | No       | Max facts to return (default 100, max 1