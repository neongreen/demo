# Demo 5: List Entries in a .zip Archive

This demo lists the names of files inside `archive.zip` without reading their contents.
A helper script `create_archive.sh` generates the zip archive used by the program.

## Libraries Used
- Standard library `archive/zip` and `fmt`.

## Run
```bash
./create_archive.sh   # create archive.zip
go run main.go
```
Expected output:
```
a.txt
b/c.txt
```
