package main

import (
    "fmt"
        "regexp"
        )

func removeVowels(s string) string {
        r := regexp.MustCompile("[aeiouAEIOU]")
            return(r.ReplaceAllString(s, ""))
}
func main() {
        fmt.Println(removeVowels("Hello WORLD"))
}
