# Demo 3: JavaScript Plugin via Otto

This demo loads a JavaScript file using the Otto VM to execute a plugin function from Go. The JavaScript function `process` simply reverses a string. Go loads the JS source, executes it in the VM, calls `process("hello")` and prints the result.

## Libraries Used
- [`github.com/robertkrimen/otto`](https://github.com/robertkrimen/otto)
- Standard library `fmt` and `io/ioutil`.

## Run
```bash
go run main.go
```
Output should be:
```
olleh
```
