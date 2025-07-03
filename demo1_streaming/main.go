package main

import (
    "encoding/json"
    "fmt"
    "os"
)

// decoder wraps json.Decoder to provide a Skip method for this demo.
type decoder struct{ *json.Decoder }

func (d *decoder) Skip() error {
    tok, err := d.Token()
    if err != nil {
        return err
    }
    if delim, ok := tok.(json.Delim); ok {
        switch delim {
        case '{':
            for d.More() {
                if _, err := d.Token(); err != nil {
                    return err
                }
                if err := d.Skip(); err != nil {
                    return err
                }
            }
            _, err = d.Token() // consume '}'
            return err
        case '[':
            for d.More() {
                if err := d.Skip(); err != nil {
                    return err
                }
            }
            _, err = d.Token() // consume ']'
            return err
        }
    }
    return nil
}

type itemID struct {
    ID int `json:"id"`
}

func main() {
    f, err := os.Open("string.json")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    dec := &decoder{json.NewDecoder(f)}

    // consume opening {
    t, err := dec.Token()
    if err != nil {
        panic(err)
    }
    if t != json.Delim('{') {
        panic("expected start object")
    }

    var sum int
    for dec.More() {
        tok, err := dec.Token()
        if err != nil {
            panic(err)
        }
        key := tok.(string)
        if key != "items" {
            if err := dec.Skip(); err != nil {
                panic(err)
            }
            continue
        }
        t, err := dec.Token()
        if err != nil {
            panic(err)
        }
        if t != json.Delim('[') {
            panic("expected array")
        }
        for dec.More() {
            var it itemID
            if err := dec.Decode(&it); err != nil {
                panic(err)
            }
            sum += it.ID
        }
        if _, err := dec.Token(); err != nil {
            panic(err)
        }
    }
    fmt.Println(sum)
}

