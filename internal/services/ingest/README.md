# Ingest Service

This service coordinates ingest workflows after raw data arrives.

Functions you will likely implement here:

- `ProcessManualPriceSubmission`
  Take raw manual input and turn it into accepted price records.
- `ProcessAIPriceSubmission`
  Take AI output and turn it into accepted price records.
- `FilterDuplicatePriceRecords`
  Prevent repeated or conflicting entries from being stored blindly.
- `ResolveSourceMetadata`
  Attach timestamps, source tags, and confidence information.
- `PersistAcceptedPrices`
  Save clean price records through the storage layer.

This service sits between raw ingest and long-term storage.
