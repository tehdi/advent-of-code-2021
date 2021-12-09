package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    displaySum := 0
    segmentIds := []string {"a", "b", "c", "d", "e", "f", "g"}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        decoder := make(map[string]string) // key = scrambled, value = reference. In other words, "this line says <key> when I want it to say <value>"
        aAndC := make([]string, 0)
        dAndG := make([]string, 0)
        parts := strings.Split(line, " | ")
        patternPart, displayPart := parts[0], parts[1]

        for _, segmentId := range segmentIds {
            // these lines are short enough that I may be able to get away with counting
            frequency := strings.Count(patternPart, segmentId)
            switch frequency {
                case 4: decoder[segmentId] = "e"
                case 6: decoder[segmentId] = "b"
                case 7: dAndG = append(dAndG, segmentId)
                case 8: aAndC = append(aAndC, segmentId)
                case 9: decoder[segmentId] = "f"
            }
        }

        patterns := strings.Fields(patternPart)
        for _, pattern := range patterns {
            if len(pattern) == 2 {
                // for the two characters in aAndC, c is in all four of these easy numbers and a is only in 7 and 8
                // hence if it's in this pattern for 1, then it's c, and the other is a
                if IsInString(pattern, aAndC[0]) {
                    decoder[aAndC[0]] = "c"
                    decoder[aAndC[1]] = "a"
                } else {
                    decoder[aAndC[0]] = "a"
                    decoder[aAndC[1]] = "c"
                }
            } else if len(pattern) == 4 {
                // for the two characters in dAndG, d is in both 4 and 8 while g is only in 8
                // hence if it's in this pattern for 4, then it's d, and the other is g
                if IsInString(pattern, dAndG[0]) {
                    decoder[dAndG[0]] = "d"
                    decoder[dAndG[1]] = "g"
                } else {
                    decoder[dAndG[0]] = "g"
                    decoder[dAndG[1]] = "d"
                }
            }
        }
        // good news! we've decoded all the letters for this entry
        // now let's try to find an easy way to turn these into the "real" numbers we're after
        // how about treating each segment pattern as a binary number instead of a visual pattern, and adding up the segments?
        mathed := map[string]int { "a": 64, "b": 32, "c": 16, "d": 8, "e": 4, "f": 2, "g": 1 }
        displayedValueBySegmentSum := map[int]int { 119: 0, 18: 1, 93: 2, 91: 3, 58: 4, 107: 5, 111: 6, 82: 7, 127: 8, 123: 9 }
        indexMultipliers := map[int]int { 0: 1000, 1: 100, 2: 10, 3: 1 }
        displayedValue := 0
        for index, activeSegments := range strings.Fields(displayPart) {
            segmentSum := 0
            for _, segment := range activeSegments {
                decoded := decoder[string(segment)]
                segmentSum += mathed[decoded]
            }
            // segmentSum now has something we can ask displayedValueBySegmentSum for to find one digit of our 4-digit display
            digit := displayedValueBySegmentSum[segmentSum]
            displayedValue += digit * indexMultipliers[index]
        }
        displaySum += displayedValue
    } // finally exiting the "for line in file" loop
    fmt.Println(displaySum)

    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
}

func IsInString(s, substr string) bool {
    return strings.Index(s, substr) != -1
}
