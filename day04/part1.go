package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
)

type boardPosition struct {
    row, column int
    diagonal string
    marked bool
}

type bingoBoard struct {
    numbers map[int]boardPosition
    rowCounts, columnCounts map[int]int
    diagonalCounts map[string]int
}

func (b *bingoBoard) MarkNumber(number int) bool {
    // "comma ok" idiom for "if key is in map"
    if boardPosition, ok := b.numbers[number]; ok {
        b.rowCounts[boardPosition.row]++
        b.columnCounts[boardPosition.column]++
        if b.rowCounts[boardPosition.row] == 5 || b.columnCounts[boardPosition.column] == 5 {
            return true // winner!
        }

        if boardPosition.diagonal != "none" {
            b.diagonalCounts[boardPosition.diagonal]++
            if b.diagonalCounts[boardPosition.diagonal] == 5 {
                return true // winner!
            }
        }
    }
    return false
}

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    values := make([]string, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        values = append(values, line)
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println(values)
    // numbers = values[0]
    // then a blank line followed by a board (5 lines),
    // then more boards that are always 5 lines, and always with blank lines between them
}

func CalculateDiagonal(row, column int) string {
    if row == column {
        return "down"
    }
    // 4 assumes 0-index. change to 6 if you 1-index
    if (row + column == 4) {
        return "up"
    }
    return "none"
}
