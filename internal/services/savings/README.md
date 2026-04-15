# Savings Service

This service should apply the savings rule consistently.

Functions you will likely implement here:

- `CalculateSavings`
  Compute how much cheaper one supermarket is than the base supermarket.
- `IsSavingsAboveThreshold`
  Check whether savings is greater than the required amount.
- `ApplySavingsThreshold`
  Filter comparison results using the savings rule.
- `SavingsPercentage`
  Return percentage savings if you want a richer explanation later.
- `ThresholdValue`
  Return the current savings threshold from configuration.

This service is small but important because your rule about savings over `100` should live in one place, not be repeated everywhere.
