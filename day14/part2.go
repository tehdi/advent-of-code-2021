package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

type Pair struct {
    first, second string
}

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    var template string
    insertionRules := make(map[string]string)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if template == "" {
            template = scanner.Text()
            continue
        }

        line := scanner.Text()
        if line == "" {
            continue
        }

        rule := strings.Split(line, " -> ")
        insertionRules[rule[0]] = rule[1]
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    polymerChain := strings.Split(template, "")
    firstElement := polymerChain[0]
    pairs := FindPairs(polymerChain)
    fmt.Println(pairs)
    steps := 40
    for i := 0; i < steps; i++ {
        fmt.Println("Step", i+1)
        pairs = ExpandPairs(pairs, insertionRules)
    }
    elementCount := CountElements(pairs)
    elementCount[firstElement]++
    fmt.Println(elementCount)

    min := 0
    max := 0
    isAnySet := false
    for _,count := range elementCount {
        if !isAnySet {
            min = count
            max = count
            isAnySet = true
        } else {
            min = FindMin(min, count)
            max = FindMax(max, count)
        }
    }
    fmt.Printf("%d - %d = %d\n", max, min, max - min)
}

func FindPairs(polymerChain []string) map[Pair]int {
    pairs := make(map[Pair]int)
    for i := 1; i < len(polymerChain); i++ {
        first := polymerChain[i-1]
        second := polymerChain[i]
        pairs[Pair { first, second }]++
    }
    return pairs
}

func ExpandPairs(pairs map[Pair]int, insertionRules map[string]string) map[Pair]int {
    newPairs := make(map[Pair]int)
    for pair,count := range pairs {
        insertion := insertionRules[pair.first + pair.second]
        newPairs[Pair{ pair.first, insertion }] += count
        newPairs[Pair{ insertion, pair.second }] += count
    }
    return newPairs
}

func CountElements(pairs map[Pair]int) map[string]int {
    elementCount := make(map[string]int)
    for pair,count := range pairs {
        // pairs overlap so if I count both elements I'll get ~2x everything
        // but not exactly 2x because one of these pairs is at the start of the chain,
        // and another is at the end,
        // and those both have 1 element each that doesn't overlap with anything else
        // so I should either count the first in each pair and treat the last element in the original template as a special case,
        // or count the second and treat the original first as a special case.
        // the first is easier to parse from the initial template
        elementCount[pair.second] += count
    }
    return elementCount
}

func FindMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func FindMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}
