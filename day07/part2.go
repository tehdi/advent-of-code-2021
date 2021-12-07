package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
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

    const crabmarineCount = 1000
    crabmarinePositions := make(map[int]int)
    crabmarinePositionSum := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for _, stringValue := range strings.Split(line, ",") {
            intValue, _:= strconv.Atoi(stringValue)
            crabmarinePositionSum += intValue
            crabmarinePositions[intValue]++
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    //fmt.Println(crabmarinePositions)
    average := crabmarinePositionSum / crabmarineCount
    fmt.Println("Target position:", average)

    fuelNeeded := 0
    for position, count := range crabmarinePositions {
        var steps int
        if position > average {
            steps = position - average
        } else {
            steps = average - position
        }
        if steps > 0 {
            thisFuel := calculateFuelNeeded(steps)
            fmt.Printf("Starting position %d: %d crabs * %d steps * %d fuel per crab\n", position, count, steps, thisFuel)
            fuelNeeded += thisFuel * count
        }
    }

    fmt.Println("Fuel needed:", fuelNeeded)
}

func calculateFuelNeeded(steps int) int {
    // (number of steps + 1) * (number of steps / 2)
    return int(float64(steps + 1) * (float64(steps) / 2.0))
}
