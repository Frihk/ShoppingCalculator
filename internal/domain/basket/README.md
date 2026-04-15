# Basket Domain

Use this area for shopping basket concepts and rules.

Functions you will likely implement here:

- `NewBasket`
  Create a new basket with a customer-selected base supermarket.
- `AddItem`
  Add an item and quantity to the basket.
- `RemoveItem`
  Remove an item from the basket.
- `UpdateItemQuantity`
  Change the quantity of an item already in the basket.
- `ListItems`
  Return all items in the basket.
- `ValidateBasket`
  Ensure the basket is usable for comparison.
- `SetBaseSupermarket`
  Define the supermarket used as the savings reference point.
- `SelectedSupermarkets`
  Return the supermarkets chosen for comparison.

Rules this domain should protect:

- quantity cannot be zero or negative
- item names should not be empty
- base supermarket should be clearly defined before comparison
