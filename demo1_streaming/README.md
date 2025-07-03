# Demo 1: Streaming Mixed Arrays and Objects

This demo shows how to use Go's `encoding/json` `Decoder` to stream through a JSON document that mixes objects and arrays. It reads `string.json`, skips unneeded fields, sums the `id` values inside the `items` array, and prints the total.

## Libraries Used
- Standard library `encoding/json`, `fmt`, and `os`.

## Run
```bash
go run main.go
```
It should output:
```
3
```
