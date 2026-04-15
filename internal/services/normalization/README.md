# Normalization Service

This service should make inconsistent names comparable.

Functions you will likely implement here:

- `NormalizeProductName`
  Standardize incoming product names.
- `NormalizeSupermarketName`
  Standardize supermarket names.
- `MatchProductAliases`
  Map alternative spellings or labels to one product identity.
- `MatchSupermarketAliases`
  Map alternative store names to one supermarket identity.
- `NormalizePriceBatch`
  Apply normalization rules to many incoming entries at once.

This becomes important as soon as user input and AI scraping use different wording.
