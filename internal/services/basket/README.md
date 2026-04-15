# Basket Service

This service coordinates basket workflows.

Functions you will likely implement here:

- `CreateBasket`
  Build a basket from validated request data.
- `AddBasketItem`
  Add an item through service-level rules.
- `UpdateBasketItem`
  Change an existing basket item.
- `RemoveBasketItem`
  Remove an item safely.
- `ValidateBasketForComparison`
  Confirm the basket is ready to be used for supermarket comparison.
- `AttachBaseSupermarket`
  Set the user-selected reference supermarket.
- `AttachComparisonSupermarkets`
  Add the supermarkets the user wants to compare.

This service may call domain objects plus repositories if baskets are persisted.
