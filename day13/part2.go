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

    var grid [655][892]bool
    xFolds := []int { 655, 327, 163, 81, 40 }
    yFolds := []int { 447, 223, 111, 55, 27, 13, 6 }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        coords := strings.Split(line, ",")
        x,_ := strconv.Atoi(coords[0])
        y,_ := strconv.Atoi(coords[1])

        for _, xFold := range xFolds {
            x = Fold(x, xFold)
        }
        for _, yFold := range yFolds {
            y = Fold(y, yFold)
        }

        grid[y][x] = true
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println("Final dot count:", CountDots(grid))
    PrintManual(grid)
}

func Fold(old, fold int) int {
    if old < fold {
        return old
    }
    overage := old - fold
    return fold - overage
}

func CountDots(grid [655][892]bool) int {
    dotCounter := 0
    for _, x := range grid {
        for _, value := range x {
            if value {
                dotCounter++
            }
        }
    }
    return dotCounter
}

func PrintManual(grid [655][892]bool) {
    for xIndex, x := range grid {
        if xIndex > 6 {
            break
        }
        fmt.Print("  ")
        for yIndex, value := range x {
            if yIndex > 40 {
                break
            }
            if value {
                fmt.Print("*")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}
