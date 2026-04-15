# Backend Structure

This project is organized around the backend domain instead of the old calculator flow.

Core direction:

- `domain` holds the business concepts
- `services` holds use-case logic
- `storage` holds persistence responsibilities
- `ingest` handles incoming price data
- `api` exposes the backend
- `config` holds environment and rules

Suggested core concepts:

- basket
- basket item
- supermarket
- product
- price record
- price source
- recommendation result
- savings threshold

Suggested recommendation modes:

- AI-based comparison
- consumer-input comparison
- baseline cheapest-supermarket comparison
