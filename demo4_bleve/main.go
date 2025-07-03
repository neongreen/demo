package main

import (
    "fmt"

    "github.com/blevesearch/bleve"
)

func main() {
    mapping := bleve.NewIndexMapping()
    index, err := bleve.New("demo.bleve", mapping)
    if err == bleve.ErrorIndexPathExists {
        index, err = bleve.Open("demo.bleve")
    }
    if err != nil {
        panic(err)
    }
    defer index.Close()

    docs := []struct{
        ID string
        Body string
    }{
        {"1", "go is great"},
        {"2", "python is popular"},
        {"3", "go and python"},
    }

    for _, d := range docs {
        if err := index.Index(d.ID, d); err != nil {
            panic(err)
        }
    }

    if err := index.Close(); err != nil {
        panic(err)
    }

    index, err = bleve.Open("demo.bleve")
    if err != nil {
        panic(err)
    }
    defer index.Close()

    query := bleve.NewMatchQuery("go")
    search := bleve.NewSearchRequest(query)
    res, err := index.Search(search)
    if err != nil {
        panic(err)
    }
    for _, hit := range res.Hits {
        fmt.Println(hit.ID)
    }
}

