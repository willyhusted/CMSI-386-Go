package main

import (
    "fmt"
    "strings"
    "math/rand"
    "math"
    "regexp"
)

func change(amt int) (int, int, int, int) {
        quarterVal := 25
        dimeVal := 10
        nickelVal := 5
        quarters := amt/quarterVal
        amt = amt%quarterVal
        dimes := amt/dimeVal
        amt = amt%dimeVal
        nickels := amt/nickelVal
        pennies := amt%nickelVal
        return quarters, dimes, nickels, pennies
}

func removeVowels(s string) string {
        r := regexp.MustCompile("[aeiouAEIOU]")
        return(r.ReplaceAllString(s, ""))
}

func scramble(s string) (string){
        a := (strings.Split("Hello World", ""))
        for i := len(a); i > 0; i--{
            j := int(math.Floor(float64(i)*rand.Float64()))
            a[i-1], a[j] = a[j], a[i-1]
        }
        return strings.Join(a, " ")
}

func powersOfTwo(limit int, callback func(int)) {
    result := 1
        for result <= limit {
            callback(result)
            result = result * 2
        }
}

func powers(base int, limit int, callback func(int)) {
    result := 1
        for result <= limit {
            callback(result)
            result = result * base
            }
}

func interleave(a []string, b []string) []string {
        alength := len(a)
        blength := len(b)
        max := math.Max(float64(alength), float64(blength))
        result := make([]string, (0))
        for i := 0; i < int(max); i+=1 {
            if i < alength {
                result = append(result, a[i])}
            if i < blength {
                result = append(result, b[i])}
        }
        return result
}

func stutter(a []string) []string{
        return interleave(a, a)
}

func unitTest(){
	
}

func main() {
        fmt.Println(change(78))
        fmt.Println(removeVowels("Hello WORLD"))
        fmt.Println(scramble("Hello World"))
//      fmt.Println(powersOfTwo())
        fmt.Println(interleave([]string{"1", "2"}, []string{"a", "b", "c"}))
        fmt.Println(stutter([]string{"4", "3"}))
}
