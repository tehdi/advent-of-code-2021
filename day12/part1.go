package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

// Node graph from https://flaviocopes.com/golang-data-structure-graph/
type (
    Cave struct {
        name string
    }
    CaveMap struct {
        caves []*Cave
        paths map[*Cave][]*Cave
    }
)

func (this *Cave) IsBig() bool {
    return this.name == strings.ToUpper(this.name)
}
func (this *Cave) IsSmall() bool {
    return this.name == strings.ToLower(this.name)
}
func (this *Cave) String() string {
    return fmt.Sprintf("%v", this.name)
}
func (this *CaveMap) AddCave(cave *Cave) {
    this.caves = append(this.caves, cave)
}
func (this *CaveMap) AddPath(c1, c2 *Cave) {
    if this.paths == nil {
        this.paths = make(map[*Cave][]*Cave)
    }
    this.paths[c1] = append(this.paths[c1], c2)
    this.paths[c2] = append(this.paths[c2], c1)
}
func (this CaveMap) String() string {
    var s strings.Builder
    for _, cave := range this.caves {
        s.WriteString(cave.String() + " -> ")
        for _, adjacent := range this.paths[cave] {
            s.WriteString(adjacent.String() + " ")
        }
        s.WriteString("\n")
    }
    return fmt.Sprintf("%s", s.String())
}

// Stack from https://github.com/golang-collections/collections/blob/604e922904d3/stack/stack.go
type (
    Stack struct {
        top *Node
        length int
    }
    Node struct {
        value *Cave
        prev *Node
    }
)
func NewStack() *Stack {
    return &Stack{nil,0}
}
func (this *Stack) Len() int {
    return this.length
}
func (this *Stack) Peek() *Cave {
    return this.top.value
}
func (this *Stack) Pop() *Cave {
    n := this.top
    this.top = n.prev
    this.length--
    return n.value
}
func (this *Stack) Push(cave *Cave) {
    n := &Node{cave,this.top}
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

    var caveMap CaveMap
    var start, end *Cave
    caves := make(map[string]*Cave)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        caveNames := strings.Split(line, "-")

        for _, caveName := range caveNames {
            // if we haven't seen this cave yet, add it to all the trackers
            if _, ok := caves[caveName]; !ok {
                cave := Cave { caveName }
                if caveName == "start" {
                    start = &cave
                } else if caveName == "end" {
                    end = &cave
                }
                caves[caveName] = &cave
                caveMap.AddCave(&cave)
            }
        }
        caveMap.AddPath(caves[caveNames[0]], caves[caveNames[1]])
    }
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }

    visitCounter := make(map[*Cave]int)
    path := NewStack()
    VisitCave(&caveMap, start, end, path, visitCounter)
    fmt.Println("Paths found to get us out of here:", visitCounter[end])
}

func VisitCave(caveMap *CaveMap, cave *Cave, end *Cave, path *Stack, visitCounter map[*Cave]int) {
    //fmt.Println("Entering", cave)
    path.Push(cave)
    visitCounter[cave]++

    if cave.name == "end" {
        // fmt.Println(" found a way out")
        path.Pop()
        return
    }
    for _, adjacent := range caveMap.paths[cave] {
        // fmt.Println(" Adjacency check:", cave, adjacent)
        if CanVisitCave(adjacent, visitCounter) {
            VisitCave(caveMap, adjacent, end, path, visitCounter)
        } else {
            // fmt.Println("  Cannot visit adjacent cave", adjacent)
        }
    }
    // exhausted all possibilities from this cave
    // time to back out and try a different path, if there are any left
    // fmt.Println("Backing out of", cave)
    path.Pop()
    visitCounter[cave]--
    return
}

func CanVisitCave(cave *Cave, visitCounter map[*Cave]int) bool {
    // I can always visit the end cave. Other than that,
    // I can visit a cave if it hasn't already been visited,
    // or if it's a big cave that I'm allowed to visit more than once.
    return cave.name == "end" || visitCounter[cave] < 1 || cave.IsBig()
}
