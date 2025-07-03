# Demo 2: Custom Timestamp Field

This demo defines a struct `Event` with custom JSON marshaling and unmarshaling so that the `Time` field is encoded using RFC3339 in UTC. A Go unit test verifies that encoding and decoding preserve the value exactly.

## Libraries Used
- Standard library `encoding/json`, `time`, and `testing`.

## Run
Execute the tests:
```bash
go test ./
```
They should pass with no output.
