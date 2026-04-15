# Ingest

This folder handles raw incoming external data before it becomes trusted internal data.

Functions you will likely implement here:

- `IngestManualPriceInput`
  Accept raw supermarket prices entered by the user.
- `IngestAIPriceInput`
  Accept raw AI-collected price results.
- `NormalizeIncomingProductNames`
  Clean incoming names before comparison.
- `TagIncomingSource`
  Attach source identity such as AI, consumer input, or baseline.
- `ValidateIncomingPriceBatch`
  Ensure the batch has enough structure to continue.
- `TransformToPriceRecords`
  Convert raw input into internal price records.
- `RejectInvalidEntries`
  Drop or report malformed data before it pollutes your system.

What does not belong here:

- final supermarket recommendation decisions
- HTTP response formatting
- long-term persistence rules
