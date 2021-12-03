package main

import (
    "fmt"
    "reflect"
)

func main() {
    const value string= "1001"
    for i := 0; i < len(value); i++ {
        fmt.Println(value[i] - 48) // this gives me the byte value, which is 0=48 1=49
        fmt.Println(string(value[i])) // this gives me 1 or 0
        fmt.Println(reflect.TypeOf(value[i])) // uint8
    }
}
