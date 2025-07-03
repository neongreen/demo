package main

import (
    "archive/zip"
    "fmt"
)

func main() {
    r, err := zip.OpenReader("archive.zip")
    if err != nil {
        panic(err)
    }
    defer r.Close()

    for _, f := range r.File {
        fmt.Println(f.Name)
    }
}

