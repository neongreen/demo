# Demo 4: Full-Text Search & Indexing with Bleve

This demo uses the Bleve library to create a simple full-text index on disk. It indexes three short documents and then searches for the term `"go"`, printing the IDs of the matching documents.

## Libraries Used
- [`github.com/blevesearch/bleve`](https://github.com/blevesearch/bleve)
- Standard library `fmt`.

## Run
```bash
go run main.go
```
The output will be:
```
1
3
```
(the order of lines may vary.)
