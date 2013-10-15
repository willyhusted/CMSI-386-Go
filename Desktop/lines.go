package main

import (
    "os"
    "fmt"
    "strings"
    "bufio"
)

func main() {
    file, _ := os.Open(os.Args[1])
    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        currentLine := strings.TrimSpace(scanner.Text())
        if currentLine != "" && !strings.HasPrefix(currentLine, "#") {
            count += 1
        }
    }
    fmt.Println(count)    
 }
