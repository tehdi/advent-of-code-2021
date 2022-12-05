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

    hexRuneToBinaryString := map[rune]string {
        48: "0000", // 0
        49: "0001", // 1
        50: "0010", // 2
        51: "0011", // 3
        52: "0100", // 4
        53: "0101", // 5
        54: "0110", // 6
        55: "0111", // 7
        56: "1000", // 8
        57: "1001", // 9
        65: "1010", // A
        66: "1011", // B
        67: "1100", // C
        68: "1101", // D
        69: "1110", // E
        70: "1111", // F
    }

    var binaryInput strings.Builder
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for _,hexRune := range line {
            binaryInput.WriteString(hexRuneToBinaryString[hexRune])
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println(binaryInput.String())
}
