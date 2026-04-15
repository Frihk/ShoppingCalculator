# Domain

This folder contains the core business concepts of the app. Domain code should describe what the system is about, not how HTTP works and not how files are stored.

Suggested subdomains:

- `basket/`
- `product/`
- `price/`
- `recommendation/`
- `source/`
- `supermarket/`

What domain functions usually do:

- create valid business objects
- protect invariants
- expose simple business rules close to the data they belong to

What domain functions usually do not do:

- parse HTTP requests
- talk to the database directly
- run large workflows across many modules
