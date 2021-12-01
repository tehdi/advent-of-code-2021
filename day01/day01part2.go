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

    var one, two, three, index, increases int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lastWindow := one + two + three
        index++
        one = two
        two = three
        three, _ = strconv.Atoi(scanner.Text())
        thisWindow := one + two + three
        fmt.Printf("New round: %d + %d + %d = %d vs %d: ", one, two, three, thisWindow, lastWindow)

        if index < 4 {
            fmt.Println("still building first windows. nothing to compare yet")
            continue
        }

        if thisWindow > lastWindow {
            fmt.Println("increased")
            increases++
        } else {
            fmt.Println("did not increase")
        }
    }
        
    if scannerError := scanner.Err(); scannerError != nil {
        fmt.Println(scannerError)
    }
        
    fmt.Printf("Found %d increases\n", increases)
}
