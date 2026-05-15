# Hello, APort! - Go Example

A minimal Go example demonstrating basic integration with the APort API.

## Prerequisites

- Go 1.21 or later
- An APort API key

## Setup

```bash
export APORT_API_KEY="your-api-key-here"
```

## Running

```bash
cd examples/go/hello-aport
go run main.go
```

## Expected Output

```
Hello, APort! 👋
Checking API health...
APort API Status: healthy (version 1.0.0)
Integration successful! ✅
```

## What This Demonstrates

1. Creating an APort client with API key authentication
2. Making a health check request to the API
3. Handling responses and errors gracefully

## Next Steps

- Explore the [APort Go SDK](https://github.com/aporthq/aport-go)
- Check out more integration examples in this repository
