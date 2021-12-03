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

    allValues := make([]string, 0)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) > 1 {
            allValues = append(allValues, line)
        }
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    var o2 string
    var co2 string
    nextO2 := allValues
    nextCO2 := allValues
    for i := 0; i < len(allValues[0]); i++ {
        if len(o2) == 0 {
            o2ones, o2zeroes := filterByValueAtPosition(i, nextO2)
            nextO2 = findNextO2(o2ones, o2zeroes)
            fmt.Printf("O2 eval: %d ones and %d zeroes\n", len(o2ones), len(o2zeroes))
            if len(nextO2) == 1 {
                o2 = nextO2[0]
                fmt.Println("Found O2: ", o2)
            }
        }

        if len(co2) == 0 {
            co2ones, co2zeroes := filterByValueAtPosition(i, nextCO2)
            nextCO2 = findNextCO2(co2ones, co2zeroes)
            fmt.Printf("C02 eval: %d ones and %d zeroes\n", len(co2ones), len(co2zeroes))
            if len(nextCO2) == 1 {
                co2 = nextCO2[0]
                fmt.Println("Found CO2: ", co2)
            }
        }

        if len(o2) > 0 && len(co2) > 0 {
            break
        }
    }
    fmt.Printf("O2 = %s; C02 = %s\n", o2, co2)

    o2Rating, _ := strconv.ParseInt(o2, 2, 0)
    co2Rating, _ := strconv.ParseInt(co2, 2, 0)
    fmt.Printf("o2=%d; co2=%d; life support=%d\n", o2Rating, co2Rating, o2Rating * co2Rating)
}

func findNextO2(ones []string, zeroes []string) []string {
    if len(ones) >= len(zeroes) {
        return ones
    } else {
        return zeroes
    }
}

func findNextCO2(ones []string, zeroes []string) []string {
    if len(zeroes) <= len(ones) {
        return zeroes
    } else {
        return ones
    }
}

const ZERO byte = 48
const ONE byte = 49

func filterByValueAtPosition(position int, numbers []string) ([]string, []string) {
    ones := make([]string, 0)
    zeroes := make([]string, 0)
    for _, number := range(numbers) {
        if number[position] == ONE {
            ones = append(ones, number)
        } else {
            zeroes = append(zeroes, number)
        }
    }
    return ones, zeroes
}
