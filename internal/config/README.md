# Config

This folder holds application settings, not business data.

Functions you will likely implement here:

- `LoadConfig`
  Read the app configuration from environment variables, files, or defaults.
- `ValidateConfig`
  Check that required values exist and are usable.
- `DefaultConfig`
  Provide sensible fallback values for local development.
- `GetServerPort`
  Return the port the API should run on.
- `GetSavingsThreshold`
  Return the minimum saving amount required for a recommendation.
- `GetStoragePath`
  Return where local data files or database config should point.
- `GetAIConfig`
  Return AI-related settings such as provider, timeout, or enable flag.

Things config can store for this project:

- server port
- savings threshold such as `100`
- AI enabled or disabled
- storage paths
- timeout values
- environment mode

Things config should not store:

- basket data
- supermarket prices
- recommendation results
- user shopping history itself
