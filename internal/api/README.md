# API

This folder is the transport layer. Its job is to receive requests, call the correct service, and return responses.

Functions you will likely implement here:

- `RegisterRoutes`
  Connect all handlers to the HTTP router.
- `HealthHandler`
  Return a simple status response so you can confirm the server is running.
- `CreateBasketHandler`
  Accept a shopping list or basket request from the client.
- `SubmitManualPricesHandler`
  Accept user-entered supermarket prices.
- `SubmitAIPricesHandler`
  Accept AI-collected prices after they are scraped or prepared elsewhere.
- `CompareSupermarketsHandler`
  Receive a request that compares a basket across selected supermarkets.
- `GetRecommendationsHandler`
  Return the three recommendation sections for a basket.
- `ParseRequestBody`
  Decode incoming JSON safely.
- `WriteJSON`
  Send structured JSON responses.
- `WriteError`
  Return consistent error responses.

Functions that do not belong here:

- computing cheapest supermarket logic
- merging price sources
- calculating savings rules
- loading directly from persistence without going through services

Rule to remember:

API functions should be thin. If a handler starts becoming long, you are probably moving business logic into the wrong folder.
