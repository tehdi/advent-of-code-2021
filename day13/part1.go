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
    foldX := 655
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
        coords := strings.Split(line, ",")
        x,_ := strconv.Atoi(coords[0])
        y,_ := strconv.Atoi(coords[1])
        fmt.Printf("%d,%d => %d,%d\n", x, y, FoldX(x, foldX), y)
        grid[FoldX(x, foldX)][y] = true
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println("Final dot count:", CountDots(grid))
}

func FoldX(oldX, foldX int) int {
    if oldX < foldX {
        return oldX
    }
    overage := oldX - foldX
    return foldX - overage
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
