package main
 
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)
 
func main() {
    file, fileError := os.Open("input")
    if fileError != nil {
        fmt.Println(fileError)
    }
    defer file.Close()

    var last, current, increases int
    first := true
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        last = current
        current, _ = strconv.Atoi(scanner.Text())
        fmt.Printf("New round: last=%d; current=%d: ", last, current)

        if first {
            first = false
            fmt.Println("first round. nothing to compare")
            continue
        }
        
        if current > last {
            fmt.Println("increased")
            increases += 1
        } else {
            fmt.Println("did not increase")
        }
    }
        
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
        
    fmt.Printf("Found %d increases\n", increases)
}
