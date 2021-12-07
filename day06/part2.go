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

    fishes := make(map[int]int)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, ",")
        for _, stringFish := range split {
            numberFish, _ := strconv.Atoi(stringFish)
            fishes[numberFish]++
        }
    }
    fmt.Println(fishes)

    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    tomorrowFishes := fishes
    for day := 0; day < 256; day++ {
        todayFishes := make(map[int]int)
        for age, cohortSize := range tomorrowFishes {
            todayFishes[age] = cohortSize
        }
        tomorrowFishes = make(map[int]int)

        for age, cohortSize := range todayFishes {
            if age == 0 {
                tomorrowFishes[6] += cohortSize
                tomorrowFishes[8] += cohortSize
                //fmt.Printf("%d fish are breeding\n", cohortSize)
                //fmt.Println(tomorrowFishes)
            } else {
                tomorrowFishes[age - 1] += cohortSize
                //fmt.Printf("%d fish are going from %d to %d\n", cohortSize, age, age - 1)
                //fmt.Println(tomorrowFishes)
            }
        }

        if day == 255 || (day + 1) % 10 == 0 {
            totalFish := 0
            for _, cohortSize := range tomorrowFishes {
                totalFish += cohortSize
            }
            fmt.Println("End of day", day + 1, totalFish, tomorrowFishes)
        }
    }
}
