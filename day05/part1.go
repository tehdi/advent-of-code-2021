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


    const MAX_INDEX = 991
    var board [MAX_INDEX][MAX_INDEX]int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        // line format: 498,436 -> 498,932
        lineSplit := strings.SplitAfterN(line, " ", 3)
        x1y1 := strings.Split(strings.Trim(lineSplit[0], " "), ",")
        x2y2 := strings.Split(strings.Trim(lineSplit[2], " "), ",")

        x1, _ := strconv.Atoi(x1y1[0])
        y1, _ := strconv.Atoi(x1y1[1])
        x2, _ := strconv.Atoi(x2y2[0])
        y2, _ := strconv.Atoi(x2y2[1])

        if x1 == x2 {
            ymin, ymax := minMax(y1, y2)
            for y := ymin; y <= ymax; y++ {
                board[x1][y]++
            }
        }

        if y1 == y2 {
            xmin, xmax := minMax(x1, x2)
            for x := xmin; x <= xmax; x++ {
                board[x][y1]++
            }
        }
    }

    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    overlapCount := 0
    for x := 0; x < MAX_INDEX; x++ {
        for y := 0; y < MAX_INDEX; y++ {
            if board[x][y] > 1 {
                overlapCount++
            }
        }
    }

    fmt.Println(overlapCount)
}

func minMax(one, two int) (int, int) {
    if one < two {
        return one, two
    } else {
        return two, one
    }
}
