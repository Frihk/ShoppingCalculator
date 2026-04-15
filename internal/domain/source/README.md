# Source Domain

Use this area for describing where prices came from.

Functions you will likely implement here:

- `NewSource`
  Create a source definition for incoming data.
- `IsAI`
  Check whether the source is AI-based.
- `IsConsumerInput`
  Check whether the source came from manual user entry.
- `IsBaseline`
  Check whether the source came from your own stored reference data.
- `Confidence`
  Return the trust level if you decide to score sources.
- `ValidateSource`
  Ensure only supported source types are accepted.

This domain exists so you do not hardcode source decisions all over the codebase.
