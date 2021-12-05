package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

type boardPosition struct {
    row, column int
    diagonal string
    marked bool
}

type bingoBoard struct {
    rows []string
    numbers map[int]boardPosition
    rowCounts, columnCounts map[int]int
    diagonalCounts map[string]int
}

func NewBoard() bingoBoard {
    board := bingoBoard {
        rows: make([]string, 0),
        numbers: make(map[int]boardPosition),
        rowCounts: make(map[int]int),
        columnCounts: make(map[int]int),
        diagonalCounts: make(map[string]int),
    }
    return board
}

func (b *bingoBoard) AddRow(row string) bool {
    b.rows = append(b.rows, row)

    //fmt.Printf("Appended row to board: %s; board now contains: %s\n", row, b.rows)
    return len(b.rows) == 5 // 5 rows = board is done
}

func (b *bingoBoard) ParseRows() {
    for rowIndex, row := range b.rows {
        // Fields "splits the string s around each instance of one or more consecutive white space characters"
        columns := strings.Fields(row)
        for columnIndex, column := range columns {
            diagonal := CalculateDiagonal(rowIndex, columnIndex)
            number, _ := strconv.Atoi(column)
            b.numbers[number] = boardPosition {
                row: rowIndex,
                column: columnIndex,
                diagonal: diagonal,
            }
        }
    }

    //fmt.Printf("Parsed rows to board. Board now contains: %s\n", b.numbers)
}

func (b *bingoBoard) MarkNumber(number int) bool {
    // "comma ok" idiom for "if key is in map"
    if boardPosition, ok := b.numbers[number]; ok {
        boardPosition.marked = true
        b.numbers[number] = boardPosition // https://stackoverflow.com/a/42716918
        //fmt.Printf("Found number %d at %+v\n", number, boardPosition)

        b.rowCounts[boardPosition.row]++
        b.columnCounts[boardPosition.column]++
        if b.rowCounts[boardPosition.row] == 5 || b.columnCounts[boardPosition.column] == 5 {
            return true // winner!
        }

        if boardPosition.diagonal != "none" {
            if boardPosition.diagonal == "both" {
                b.diagonalCounts["up"]++
                b.diagonalCounts["down"]++
            } else {
                b.diagonalCounts[boardPosition.diagonal]++
            }
            if b.diagonalCounts["up"] == 5 || b.diagonalCounts["down"] == 5 {
                return true // winner!
            }
        }
    }
    return false
}

func (b *bingoBoard) SumUnmarkedNumbers() int {
    sum := 0
    for number, boardPosition := range(b.numbers) {
        if !boardPosition.marked {
            sum += number
        }
    }
    return sum
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

    //fmt.Println(values)
    // boards are always 5 lines, and there's always a blank line between them
    boards := make([]bingoBoard, 0)
    for _, line := range values[1:] {
        if len(line) == 0 {
            boards = append(boards, NewBoard())
        } else {
            boardBuilt := boards[len(boards) - 1].AddRow(line)
            if boardBuilt {
                //fmt.Printf("Board built: %+v\n", boards[len(boards) - 1])
                boards[len(boards) - 1].ParseRows()
            }
        }
    }
    //fmt.Printf("Built %d boards\n", len(boards))

    // numbers to be called
    numbers := make([]int, 0)
    for _, value := range strings.Split(values[0], ",") {
        number, _ := strconv.Atoi(value)
        numbers = append(numbers, number)
    }
    //fmt.Println(numbers)

    for _, number := range numbers {
        win := false
        for _, board := range boards {
            win = board.MarkNumber(number)
            if win {
                unmarkedSum := board.SumUnmarkedNumbers()
                fmt.Printf("Winner on number %d with unmarked sum %d => score %d!\n", number, unmarkedSum, number * unmarkedSum)
                break
            }
        }
        if win {
            break
        }
    }
}

func CalculateDiagonal(row, column int) string {
    // 4 assumes 0-index. change to 6 if you 1-index
    if row == column && row + column == 4 {
        return "both"
    }
    if row == column {
        return "down"
    }
    if (row + column == 4) {
        return "up"
    }
    return "none"
}
