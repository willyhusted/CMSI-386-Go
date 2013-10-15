package main

import (
    "os"
    "fmt"
    "strings"
    "bufio"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("error: file does not exist")
        os.Exit(1)
    } else {
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
 }
