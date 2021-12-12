package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "sort"
)

// Stack from https://github.com/golang-collections/collections/blob/604e922904d3/stack/stack.go
type (
    Stack struct {
        top *node
        length int
    }
    node struct {
        value string
        prev *node
    }
)
// Create a new stack
func NewStack() *Stack {
    return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
    return this.length
}
// View the top item on the stack
func (this *Stack) Peek() string {
    return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() string {
    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value string) {
    n := &node{value,this.top}
    this.top = n
    this.length++
}

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    lineCompletionScores := make([]int, 0)
    scores := map[string]int { ")": 1, "]": 2, "}": 3, ">": 4 }
    closersToOpeners := map[string]string { ")": "(", "]": "[", "}": "{", ">": "<" }
    openersToClosers := map[string]string { "(": ")", "[": "]", "{": "}", "<": ">" }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println("  *  *  *")
        //fmt.Println(line)

        openers := NewStack()
        corrupt := false
        for _, value := range line {
            token := string(value)
            //fmt.Println(token)
            if opener, isCloser := closersToOpeners[token]; isCloser {
                //fmt.Println("is closer")
                // this token is a closer.
                // if I don't have any active openers, or this isn't the closer for the latest opener,
                // then this line is corrupted
                if openers.Len() < 1 || opener != openers.Peek() {
                    //fmt.Println("closer does not match opener. line is corrupted")
                    corrupt = true
                    break
                } else {
                    // not only did I have an opener, but this was the correct closer for it! yay!
                    //fmt.Println("matches opener. popping")
                    openers.Pop()
                }
            } else {
                // token is not a closer. assumed to be an opener
                //fmt.Println("is not a closer. pushing")
                openers.Push(token)
            }
        }
        if !corrupt && openers.Len() > 0 {
            lineCompletionScore := 0
            // there were characters left after matching up all the opener-closer pairs for non-corrupt lines
            for {
                if openers.Len() < 1 {
                    break
                }
                opener := openers.Pop()
                closer := openersToClosers[opener]
                lineCompletionScore = (lineCompletionScore * 5) + scores[closer]
            }
            lineCompletionScores = append(lineCompletionScores, lineCompletionScore)
        }
    } // end of "for line in file" loop
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    sort.Ints(lineCompletionScores)
    // I was promised "there will always be an odd number of scores to consider"
    // and [odd number] / 2 rounds down (or more likely truncates) which is what I need here
    midPoint := len(lineCompletionScores) / 2
    median := lineCompletionScores[midPoint]
    fmt.Println("Middle score:", median)
}
