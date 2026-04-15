# Supermarket Domain

Use this area for store identity and supermarket-specific rules.

Functions you will likely implement here:

- `NewSupermarket`
  Create a supermarket entry.
- `ValidateSupermarketName`
  Ensure the name is present and usable.
- `NormalizeSupermarketName`
  Standardize names so duplicates are reduced.
- `IsCustom`
  Check whether the supermarket was added by the user.
- `MatchesName`
  Match incoming names to known supermarkets.
- `AvailableForComparison`
  Check whether the supermarket has enough data to be used in recommendation logic.

This domain matters because your app allows both known supermarkets and user-added supermarkets.
