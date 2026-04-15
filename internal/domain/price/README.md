# Price Domain

Use this area for price records and price-related rules.

Functions you will likely implement here:

- `NewPriceRecord`
  Create a price entry for a product at a supermarket from a given source.
- `ValidatePriceRecord`
  Ensure the amount, supermarket, product, and source are valid.
- `IsExpired`
  Decide whether a price is too old to trust.
- `MatchesProduct`
  Check whether a price belongs to a normalized product.
- `BelongsToSupermarket`
  Check whether a record belongs to a specific supermarket.
- `SourceType`
  Return whether the record came from AI, consumer input, or baseline data.
- `CollectedAt`
  Return the timestamp of the price record.

Important rule:

A price is not just an amount. It should always be tied to a product, a supermarket, a source, and a time.
