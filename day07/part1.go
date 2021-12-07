package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    crabmarinePositions := make([]int, 1000)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for index, stringValue := range strings.Split(line, ",") {
            crabmarinePositions[index], _ = strconv.Atoi(stringValue)
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    sort.Ints(crabmarinePositions)
    //fmt.Println(crabmarinePositions)
    midPoint := len(crabmarinePositions) / 2
    middleTwo := []int { crabmarinePositions[midPoint-1], crabmarinePositions[midPoint] }
    fmt.Println("Middle two:", middleTwo)
    median := (middleTwo[0] + middleTwo[1]) / 2
    fmt.Println("Median:", median)

    // now how much fuel to get everything there?
    fuelNeeded := 0
    for _, crabmarinePosition := range crabmarinePositions {
        if crabmarinePosition > median {
            fuelNeeded += crabmarinePosition - median
        } else {
            fuelNeeded += median - crabmarinePosition
        }
    }
    fmt.Println("Fuel needed to get to median:", fuelNeeded)
}
