package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

type (
    List struct {
        first *Node
        last *Node
    }
    Node struct {
        value string
        next *Node
    }
)
func NewList(template string) *List {
    this := &List { nil, nil }
    for _, element := range template {
        n := &Node { string(element), nil }
        if this.first == nil {
            // first element in the template
            this.first = n
        } else if this.last == nil {
            // second element in the template
            this.first.next = n
            this.last = n
        } else {
            // all the other elements
            this.last.next = n
            this.last = n
        }
    }
    return this
}
func (this *List) InsertAfter(value string, previous *Node) *Node {
    n := &Node { value, nil }
    oldNext := previous.next
    previous.next = n
    n.next = oldNext
    if oldNext == nil {
        this.last = n
    }
    return n
}
func (this *List) Append(value string) {
    n := &Node { value, nil }
    if this.last == nil {
        this.first.next = n
        this.last = n
    } else {
        this.last.next = n
        this.last = n
    }
}
func (this *List) String() string {
    var polymerChain strings.Builder
    for n := this.first; n.HasNext(); n = n.next {
        polymerChain.WriteString(n.value)
    }
    polymerChain.WriteString(this.last.value)
    return polymerChain.String()
}
func (this *Node) HasNext() bool {
    return this.next != nil
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

    polymerChain := NewList(template)
    // fmt.Println(insertionRules)
    // fmt.Println(template)

    // fmt.Println("Step 0:", polymerChain)
    for i := 0; i < 10; i++ {
        polymerChain = Foo(polymerChain, insertionRules, i+1)
    }

    elementFrequency := make(map[string]int)
    for _, element := range polymerChain.String() {
        elementFrequency[string(element)]++
    }
    fmt.Println(elementFrequency)
}

func Foo(polymerChain *List, insertionRules map[string]string, step int) *List {
    nextChain := NewList(polymerChain.first.value)
    for n1 := polymerChain.first; n1.HasNext(); n1 = n1.next {
        n2 := n1.next
        pair := n1.value + n2.value
        insertion := insertionRules[pair]
        // fmt.Printf("%s -> %s\n", pair, insertion)
        nextChain.Append(insertion)
        nextChain.Append(n2.value)
    }
    // these get real long. maybe don't print them
    // fmt.Printf("Step %d: %s\n", step, nextChain)
    return nextChain
}
