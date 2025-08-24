# Slogotel

A simple OpenTelemetry logging demo that shows how to send structured logs to an OTLP endpoint.

## What it does

This application:
- Connects to an OpenTelemetry log collector via OTLP (HTTP or gRPC)
- Sends structured log messages
- Runs until interrupted with Ctrl+C
- Handles connection errors gracefully by continuing to run

## Usage

### Basic usage
```bash
go run .
```

### Configuration

Set environment variables to control behavior:

- `OTEL_EXPORTER_OTLP_INSECURE=true` - Use insecure connection (no TLS)
- `OTEL_EXPORTER_OTLP_PROTOCOL=http` - Use HTTP transport (default is gRPC) 
- `OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4318` - Custom OTLP endpoint

### Examples

```bash
# Use HTTP with insecure connection
OTEL_EXPORTER_OTLP_INSECURE=true OTEL_EXPORTER_OTLP_PROTOCOL=http go run .

# Connect to custom endpoint
OTEL_EXPORTER_OTLP_ENDPOINT=http://jaeger:4318 go run .
```

## Development

```bash
# Run tests
make test

# Build binary
make build

# Clean build artifacts
make clean
```