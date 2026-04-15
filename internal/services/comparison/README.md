# Comparison Service

This service should perform the actual supermarket-to-supermarket basket comparison.

Functions you will likely implement here:

- `CompareBasketAcrossSupermarkets`
  Calculate basket totals for every selected supermarket.
- `CalculateBasketTotalForSupermarket`
  Compute the full basket cost in one supermarket.
- `BuildPerItemComparison`
  Produce item-by-item price differences across supermarkets.
- `CompareAgainstBaseSupermarket`
  Measure each supermarket against the chosen reference supermarket.
- `FindCheapestSupermarket`
  Return the lowest-cost supermarket for a given price source mode.
- `ListMissingPrices`
  Show which basket items do not have usable prices in a supermarket.

This is one of the most important services in your project. Keep it focused and testable.
