package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "math"
    "strings"
    "strconv"
)

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    var entryCount int
    values := make(map[int]int)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        entryCount++

        for i := 0; i < len(line); i++ {
            values[i] += int(line[i] - 48) // line[i] is the byte value, which is 0=48 1=49, but I just want 0 or 1
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println(values)

    // this works regardless of whether total entryCount is even or odd
    majorityCount := int(math.Floor(float64(entryCount) / 2) + 1)
    var gammaRateBuilder, epsilonRateBuilder strings.Builder
    for position := 0; position < len(values); position++ {
        sum := values[position]
        if sum >= majorityCount {
            gammaRateBuilder.WriteString("1")
            epsilonRateBuilder.WriteString("0")
        } else {
            gammaRateBuilder.WriteString("0")
            epsilonRateBuilder.WriteString("1")
        }
    }

    gammaRate, _ := strconv.ParseInt(gammaRateBuilder.String(), 2, 0)
    epsilonRate, _ := strconv.ParseInt(epsilonRateBuilder.String(), 2, 0)
    fmt.Printf("gamma=%d; epsilon=%d; power=%d\n", gammaRate, epsilonRate, gammaRate * epsilonRate)
}
