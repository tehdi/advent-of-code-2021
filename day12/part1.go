package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
)

// https://flaviocopes.com/golang-data-structure-graph/

type cave struct {
    name string
}

func (this *cave) String() string {
    return fmt.Sprintf("%v", this.name)
}

type caveMap struct {
    caves []*cave
    paths map[cave][]*cave
}

func (this *caveMap) AddCave(cave *cave) {
    this.caves = append(this.caves, cave)
}

func (this *caveMap) AddPath(c1, c2 *cave) {
    if this.paths == nil {
        this.paths = make(map[cave][]*cave)
    }
    this.paths[*c1] = append(this.paths[*c1], c2)
    this.paths[*c2] = append(this.paths[*c2], c1)
}

func (this caveMap) String() string {
    var s strings.Builder
    for _, cave := range this.caves {
        s.WriteString(cave.String() + " -> ")
        for _, adjacent := range this.paths[*cave] {
            s.WriteString(adjacent.String() + " ")
        }
        s.WriteString("\n")
    }
    return fmt.Sprintf("%s", s.String())
}

func main() {
    inputFile := flag.String("input-file", "", "file to use as input")
    flag.Parse()

    file, fileError := os.Open(*inputFile)
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    var caveMap caveMap
    caves := make(map[string]*cave)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        caveNames := strings.Split(line, "-")

        for _, caveName := range caveNames {
            // if we haven't seen this cave yet,
            // add it to the map
            if _, ok := caves[caveName]; !ok {
                cave := cave { caveName }
                caves[caveName] = &cave
                caveMap.AddCave(&cave)
            }
        }
        caveMap.AddPath(caves[caveNames[0]], caves[caveNames[1]])
    }
    fmt.Println(caveMap)

    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
}
