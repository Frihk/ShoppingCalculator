# Price Storage

Use this area for price record persistence.

Functions you will likely implement here:

- `SavePriceRecord`
  Store one accepted price record.
- `SavePriceBatch`
  Store many price records together.
- `GetPricesByProduct`
  Load price records for one product.
- `GetPricesBySupermarket`
  Load all prices for one supermarket.
- `GetPricesBySource`
  Load prices from AI, manual input, or baseline source.
- `GetComparablePrices`
  Load the prices needed to compare a basket across supermarkets.
- `DeleteExpiredPrices`
  Clean out stale entries if you apply freshness rules.

This repository is central to the whole app.
