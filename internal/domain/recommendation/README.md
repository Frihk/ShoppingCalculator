# Recommendation Domain

Use this area for recommendation result types and recommendation rules.

Functions you will likely implement here:

- `NewRecommendation`
  Create a recommendation result for one supermarket comparison.
- `SavingsAmount`
  Return how much money is saved compared with the base supermarket.
- `IsQualified`
  Return whether the recommendation passes the savings threshold.
- `RecommendationMode`
  Return whether the result is AI-based, consumer-input-based, or baseline-based.
- `BestSupermarket`
  Return the winning supermarket for a given mode.
- `PriceBreakdown`
  Return item-by-item comparison details.

Important rule:

This domain should describe what a recommendation means. It should not fetch prices by itself.
