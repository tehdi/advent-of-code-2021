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

    var fishes string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fishes = scanner.Text()
    }

    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fishes = strings.ReplaceAll(fishes, ",", "")
    var tomorrowFishes strings.Builder
    tomorrowFishes.WriteString(fishes)
    for day := 0; day < 80; day++ {
        // reset for the new day
        todayFishes := tomorrowFishes.String()
        tomorrowFishes.Reset()

        fmt.Printf("Starting day %d with %d fish\n", day + 1, len(todayFishes))
        for _, stringFish := range strings.SplitAfter(todayFishes, "") {
            fish, _ := strconv.Atoi(stringFish)
            if fish == 0 {
                //fmt.Printf("%d: 0 => 6+8\n", fish)
                tomorrowFishes.WriteString("6") // 0 becomes 6
                tomorrowFishes.WriteString("8")
            } else {
                //fmt.Printf("%d: other => %d\n", fish, fish - 1)
                tomorrowFishes.WriteString(strconv.Itoa(fish - 1)) // each other number decreases by 1
            }
        }

        finalFishes := tomorrowFishes.String()
        fmt.Printf("Ending day %d with %d fish\n", day + 1, len(finalFishes))
    }
}
