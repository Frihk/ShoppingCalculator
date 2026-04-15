# Recommendation Service

This service builds the final recommendation output shown to the user.

Functions you will likely implement here:

- `BuildRecommendations`
  Create the full recommendation response.
- `BuildAIRecommendations`
  Build recommendations using AI-collected prices only.
- `BuildConsumerInputRecommendations`
  Build recommendations using manual consumer-entered prices only.
- `BuildBaselineRecommendations`
  Build recommendations using your basic stored price set.
- `SelectBestQualifiedSupermarket`
  Return the best supermarket that passes the savings rule.
- `AssembleRecommendationSections`
  Format the three sections into one response structure.
- `ExplainRecommendation`
  Generate the reasons behind a recommendation, such as savings and price differences.

Important rule:

This service should use the comparison and savings services. It should not duplicate their logic carelessly.
