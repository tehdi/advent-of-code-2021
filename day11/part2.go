package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strconv"
)

const minRowIndex = 0
const maxRowIndex = 9
const minColumnIndex = 0
const maxColumnIndex = 9

type octopus struct {
    row, column int
    energyLevel int
    flashed bool
}

func NewOctopus(row, column, energyLevel int) *octopus {
    octopus := octopus {
        row: row,
        column: column,
        energyLevel: energyLevel,
        flashed: false,
    }
    return &octopus
}

func (this *octopus) Step() {
    this.energyLevel++
    if this.energyLevel > 9 {
        this.flashed = true
        this.energyLevel = 0
    }
}

func (this *octopus) String() string {
    flash := " "
    if this.energyLevel > 9 {
        flash = "*"
    }
    return fmt.Sprintf("(%d,%d: %2d %s) ", this.row, this.column, this.energyLevel, flash)
}

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    row := 0
    var octopodes [10][10]*octopus
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        for column, value := range line {
            energyLevel, _ := strconv.Atoi(string(value))
            octopodes[row][column] = NewOctopus(row, column, energyLevel)
        }
        row++
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    PrintOctopodes(octopodes, 0)
    for step := 0; ; step++ {
        stepFlashCount := 0
        // increment everyone's energy level and see who, if anyone, flashes this step
        flashers := make([]*octopus, 0)
        for rowIndex, row := range octopodes {
            for columnIndex, _ := range row {
                octopus := octopodes[rowIndex][columnIndex]
                octopus.flashed = false // reset before stepping
                octopus.Step()
                if octopus.flashed {
                    //fmt.Println("Flashed:", octopus)
                    flashers = append(flashers, octopus)
                }
            }
        }

        // increment the buddies adjacent to the primary flashers and see if anyone flashes ("secondary flashers")
        // repeat this until nobody flashes
        stepFlashCount += len(flashers)
        nextFlashers := Foo(flashers, octopodes)
        stepFlashCount += len(nextFlashers)
        for len(nextFlashers) > 0 {
            nextFlashers = Foo(nextFlashers, octopodes)
            stepFlashCount += len(nextFlashers)
        }

        if stepFlashCount >= 100 {
            fmt.Printf("%d octopodes flashed on step %d\n", stepFlashCount, step + 1)
            PrintOctopodes(octopodes, step + 1)
            break
        }
        if (step+1) % 10 == 0 {
            PrintOctopodes(octopodes, step+1)
        }
    }
    PrintOctopodes(octopodes, 0)
}

func PrintOctopodes(octopodes [10][10]*octopus, step int) {
    fmt.Println("Step:", step)
    for _, row := range octopodes {
        for _, octopus := range row {
            fmt.Print(octopus)
        }
        fmt.Println()
    }
    fmt.Println()
}

func Foo(flashers []*octopus, octopodes [10][10]*octopus) []*octopus {
    nextFlashers := make([]*octopus, 0)
    for _, flasher := range flashers {
        // fmt.Println("Finding adjacent", flasher)
        validRows := make([]int, 1)
        validColumns := make([]int, 1)
        validRows[0] = flasher.row
        validColumns[0] = flasher.column
        AppendPrevious(&validRows, flasher.row, minRowIndex)
        AppendNext(&validRows, flasher.row, maxRowIndex)
        AppendPrevious(&validColumns, flasher.column, minColumnIndex)
        AppendNext(&validColumns, flasher.column, maxColumnIndex)
        // fmt.Println("Found:", validRows, validColumns)
        for _, validRow := range validRows {
            for _, validColumn := range validColumns {
                if validRow != flasher.row || validColumn != flasher.column {
                    // if one octopus is adjacent to multiple flashers, it'll make it here multiple times
                    // and will be incremented multiple times. exactly what we want
                    adjacentBuddy := octopodes[validRow][validColumn]
                    if !adjacentBuddy.flashed {
                        adjacentBuddy.Step()
                        if adjacentBuddy.flashed {
                            nextFlashers = append(nextFlashers, adjacentBuddy)
                        }
                    }
                }
            }
        }
    }
    return nextFlashers
}

func AppendPrevious(validIndexes *[]int, thisIndex, minIndex int) {
    if thisIndex > minIndex {
        *validIndexes = append(*validIndexes, thisIndex - 1)
    }
}

func AppendNext(validIndexes *[]int, thisIndex, maxIndex int) {
    if thisIndex < maxIndex {
        *validIndexes = append(*validIndexes, thisIndex + 1)
    }
}
