package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
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

    // does looking at the input for limits count as cheating?
    var heightmap [100][100]int
    activeRow := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for index, value := range line {
            height, _ := strconv.Atoi(string(value))
            heightmap[activeRow][index] = height
        }
        activeRow++
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    minRowIndex := 0
    maxRowIndex := len(heightmap) - 1
    minColumnIndex := 0
    maxColumnIndex := len(heightmap[0]) - 1
    riskLevel := 0
    for rowIndex, row := range heightmap {
        for columnIndex, value := range row {
            lowPoint := true
            if rowIndex > minRowIndex {
                upValue := heightmap[rowIndex-1][columnIndex]
                lowPoint = lowPoint && value < upValue
            }
            if rowIndex < maxRowIndex {
                downValue := heightmap[rowIndex+1][columnIndex]
                lowPoint = lowPoint && value < downValue
            }
            if columnIndex > minColumnIndex {
                leftValue := heightmap[rowIndex][columnIndex-1]
                lowPoint = lowPoint && value < leftValue
            }
            if columnIndex < maxColumnIndex {
                rightValue := heightmap[rowIndex][columnIndex+1]
                lowPoint = lowPoint && value < rightValue
            }
            if lowPoint {
                riskLevel += value + 1
            }
        }
    }
    fmt.Println(riskLevel)
}
