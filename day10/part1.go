package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
)

// Stack from https://github.com/golang-collections/collections/blob/604e922904d3/stack/stack.go
type (
    Stack struct {
        top *node
        length int
    }
    node struct {
        value interface{}
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
func (this *Stack) Peek() interface{} {
    if this.length == 0 {
        return nil
    }
    return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
    if this.length == 0 {
        return nil
    }
    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
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

    syntaxErrorScore := 0
    scores := map[string]int { ")": 3, "]": 57, "}": 1197, ">": 25137 }
    closersToOpeners := map[string]string { ")": "(", "]": "[", "}": "{", ">": "<" }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        //fmt.Println("  *  *  *")
        //fmt.Println(line)

        openers := NewStack()
        for _, value := range line {
            token := string(value)
            //fmt.Println(token)
            if opener, isCloser := closersToOpeners[token]; isCloser {
                //fmt.Println("is closer")
                // this token is a closer.
                // if I don't have any active openers, or this isn't the closer for the latest opener,
                // then this line is corrupted
                if openers.Len() < 1 || opener != openers.Peek() {
                    //fmt.Println("does not match opener. incrementing score")
                    syntaxErrorScore += scores[token]
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
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    fmt.Println(syntaxErrorScore)
}
