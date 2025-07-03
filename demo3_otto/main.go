package main

import (
    "fmt"
    "io/ioutil"

    "github.com/robertkrimen/otto"
)

func main() {
    src, err := ioutil.ReadFile("plugin.js")
    if err != nil {
        panic(err)
    }

    vm := otto.New()
    if _, err := vm.Run(src); err != nil {
        panic(err)
    }

    value, err := vm.Call("process", nil, "hello")
    if err != nil {
        panic(err)
    }
    result, err := value.ToString()
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}

