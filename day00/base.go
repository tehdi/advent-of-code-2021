package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
)

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
}
