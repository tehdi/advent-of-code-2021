package main
 
import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strconv"
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

    var horizontal, depth int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " ")

        direction := parts[0]
        value, _ := strconv.Atoi(parts[1])

        switch direction {
            case "forward":
                horizontal += value
            case "down":
                depth += value
            case "up":
                depth -= value
        }
        fmt.Printf("Moving %s %d => h:%d * d:%d = %d\n", direction, value, horizontal, depth, horizontal * depth)
    }
        
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
}
