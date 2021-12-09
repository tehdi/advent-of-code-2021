package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    easyNumberCount := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // when parsing a line ignore everything before the |
        // then split the remaining values
        // and count how many have length = { 2, 3, 4, 7 }
        displays := strings.Fields(strings.Split(line, " | ")[1])
        for _, display := range displays {
            if hasUniqueLength(display) {
                easyNumberCount++
            }
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println(easyNumberCount)
}

func hasUniqueLength(display string) bool {
    length := len(display)
    return length == 2 || length == 3 || length == 4 || length == 7
}
