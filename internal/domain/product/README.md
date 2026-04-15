# Product Domain

Use this area for product identity and naming rules.

Functions you will likely implement here:

- `NewProduct`
  Create a canonical product definition.
- `NormalizeName`
  Standardize item names so similar entries can be compared.
- `ValidateProductName`
  Reject empty or unusable names.
- `MatchesInputName`
  Check whether a raw incoming name maps to a stored product.
- `Category`
  Return the product category if you decide to support one.
- `Unit`
  Return the expected measurement unit if you later support sizes.

This domain matters because AI input and user input will rarely use perfectly consistent naming.
