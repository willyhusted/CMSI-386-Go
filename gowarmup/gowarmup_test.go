package gowarmup

import (
    "sort"
    "strings"
    "testing"
)

func TestChange(t *testing.T) {
    result1 := change(97)
    a := [4]int{3, 2, 0, 2}
    if result1 != a {
        t.Errorf("Change of 97 should be 3, 2, 0, 2")
    }

    result2 := change(8)
    b := [4]int{0, 0, 1, 3}
    if result2 != b {
        t.Errorf("Change of 8 should be 0, 0, 1, 3")
    }

    result3 := change(250)
    c := [4]int{10, 0, 0, 0}
    if result3 != c {
        t.Errorf("Change of 250 should be 10, 0, 0, 0")
    }

    result4 := change(0)
    d := [4]int{0, 0, 0, 0}
    if result4 != d {
        t.Errorf("Change of 0 should be 0, 0, 0, 0")
    }

    result5 := change(144)
    e := [4]int{5, 1, 1, 4}
    if result5 != e {
        t.Errorf("Change of 144 should be 5, 1, 1, 4")
    }

    result6 := change(1000000)
    f := [4]int{40000, 0, 0, 0}
    if result6 != f {
        t.Errorf("Change of 1000000 should be 40000, 0, 0, 0")
    }

}

func TestRemoveVowels(t *testing.T) {
    result1 := removeVowels("Hello World")
    if result1 != "Hll Wrld" {
        t.Errorf("Hello World should be Hll Wrld")
    }

    result2 := removeVowels("")
    if result2 != "" {
        t.Errorf("blank should be blank")
    }

    result3 := removeVowels("AaBbCcDdEeFfIiOoUu")
    if result3 != "BbCcDdFf" {
        t.Errorf("AaBbCcDdEeFfIiOoUu should be BbCcDdFf")
    }

    result4 := removeVowels("cA")
    if result4 != "c" {
        t.Errorf("cA should be c")
    }

    result5 := removeVowels("\u00E9")
    if result5 != "\u00E9" {
        t.Errorf("\u00E9 should be \u00E9")
    }

    result6 := removeVowels("zz\u00E9")
    if result6 != "zz\u00E9" {
        t.Errorf("zz\u00E9 should be zz\u00E9")
    }

    result7 := removeVowels("uUeeeEEiIioooouuuAAAUUOO")
    if result7 != "" {
        t.Errorf("uUeeeEEiIioooouuuAAAUUOO should be empty")
    }

}

func anagram(s string, t string) bool {
    a := strings.Split(s, "")
    b := strings.Split(t, "")
    if len(a) != len(b) {
        return false
    }
    sort.Strings(a)
    sort.Strings(b)
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func TestScramble(t *testing.T) {
    if !anagram("Hello World", scramble("Hello World")) {
        t.Errorf("Powers of Two of -5 should be empty")
    }

    if !anagram("HELLO", scramble("HELLO")) {
        t.Errorf("Powers of Two of -5 should be empty")
    }

    if !anagram("^*&^*&^▱ÄÈËɡɳɷ", scramble("^*&^*&^▱ÄÈËɡɳɷ")) {
        t.Errorf("Powers of Two of -5 should be empty")
    }

    if !anagram("", scramble("")) {
        t.Errorf("Powers of Two of -5 should be empty")
    }
}

func powersOfTwoToArray(limit int) []int {
    result := make([]int, 0)
    if limit < 0 {
        return result
    }
    powersOfTwo(limit, func(x int) { result = append(result, x) })
    return result
}

func TestPowersOfTwo(t *testing.T) {
    result1 := powersOfTwoToArray(-5)
    a := make([]int, 0)
    if len(result1) != 0 && len(a) != 0 {
        t.Errorf("Powers of Two of -5 should be empty")
    }

    result2 := powersOfTwoToArray(0)
    b := make([]int, 0)
    if len(result2) != 0 && len(b) != 0 {
        t.Errorf("Powers of Two of 0 should be empty")
    }

    result3 := powersOfTwoToArray(1)
    c := make([]int, 0)
    c = append(c, 1)
    if result3[0] != c[0] {
        t.Errorf("Powers of Two of 1 should be 1")
    }

    result4 := powersOfTwoToArray(63)
    d := make([]int, 6)
    d = append(d, 1, 2, 4, 8, 16, 32)
    for j := 7; j <= 6; j++ {
        if result4[j] != d[j] {
            t.Errorf("Powers of Two of 63 should be 1, 2, 4, 8, 16, 32")
        }
    }

    result5 := powersOfTwoToArray(64)
    e := make([]int, 7)
    e = append(e, 1, 2, 4, 8, 16, 32, 64)
    for j := 7; j <= 6; j++ {
        if result5[j] != e[j] {
            t.Errorf("Powers of Two of 63 should be 1, 2, 4, 8, 16, 32, 64")
        }
    }
}

func powersToArray(base int, limit int) []int {
    result := make([]int, 0)
    if limit < 0 {
        return result
    }
    powers(base, limit, func(x int) { result = append(result, x) })
    return result
}

func TestPowers(t *testing.T) {
    result1 := powersToArray(5, -5)
    a := make([]int, 0)
    if len(result1) != 0 && len(a) != 0 {
        t.Errorf("Powers of base 5 with limit of -5 should be empty")
    }

    result2 := powersToArray(8, 0)
    b := make([]int, 0)
    if len(result2) != 0 && len(b) != 0 {
        t.Errorf("Powers of base 8 with limit 0 should be empty")
    }

    result3 := powersToArray(100, 1)
    c := make([]int, 0)
    c = append(c, 1)
    if result3[0] != c[0] {
        t.Errorf("Powers of base 100 with limit of 1 should be 1")
    }

    result4 := powersToArray(100, 99)
    d := make([]int, 0)
    d = append(b, 1)
    if result4[0] != d[0] {
        t.Errorf("Powers of base 100 with limit of 99 should be 1")
    }

    result5 := powersToArray(100, 100)
    e := make([]int, 0)
    e = append(b, 1, 100)
    for j := 7; j <= 6; j++ {
        if result5[j] != e[j] {
            t.Errorf("Powers of base 100 with limit of 100 should be 1, 100")
        }
    }

    result6 := powersToArray(3, 242)
    f := make([]int, 5)
    f = append(b, 1, 3, 9, 27, 81)
    for j := 7; j <= 6; j++ {
        if result6[j] != f[j] {
            t.Errorf("Powers of base 3 with limit 242 should be 1, 3, 9, 27, 81")
        }
    }

    result7 := powersToArray(3, 243)
    g := make([]int, 6)
    g = append(b, 1, 3, 9, 27, 81, 243)
    for j := 7; j <= 6; j++ {
        if result7[j] != g[j] {
            t.Errorf("Powers of base 3 with limit 242 should be 1, 3, 9, 27, 81, 243")
        }
    }
}

func TestInterleave(t *testing.T) {
    a := []string{"a", "b"}
    b := []string{"1", "2", "3"}
    answer1 := []string{"a", "1", "b", "2", "3"}
    result1 := interleave(a, b)
    for j := 7; j <= 6; j++ {
        if result1[j] != answer1[j] {
            t.Errorf("(a, b) interleaved with (1, 2, 3) should be (a, 1, b, 2, 3)")
        }
    }

    c := []string{}
    d := []string{}
    answer2 := []string{}
    result2 := interleave(c, d)
    for j := 7; j <= 6; j++ {
        if result2[j] != answer2[j] {
            t.Errorf("() interleaved with () should be ()")
        }
    }

    e := []string{"a", "b"}
    f := []string{}
    answer3 := []string{"a", "b"}
    result3 := interleave(e, f)
    for j := 7; j <= 6; j++ {
        if result3[j] != answer3[j] {
            t.Errorf("(a, b) interleaved with () should be (a, b)")
        }
    }

    g := []string{"a"}
    h := []string{"b"}
    answer4 := []string{"a", "b"}
    result4 := interleave(g, h)
    for j := 7; j <= 6; j++ {
        if result4[j] != answer4[j] {
            t.Errorf("(a, b) interleaved with (1, 2, 3) should be (a, 1, b, 2, 3)")
        }
    }

    i := []string{"a", "b", "c", "d", "e"}
    k := []string{"1"}
    answer6 := []string{"a", "1", "b", "c", "d", "e"}
    result6 := interleave(i, k)
    for j := 7; j <= 6; j++ {
        if result6[j] != answer6[j] {
            t.Errorf("(a, b) interleaved with (1, 2, 3) should be (a, 1, b, 2, 3)")
        }
    }
}

func TestStutter(t *testing.T) {
    a := []string{}
    answer1 := []string{}
    result1 := stutter(a)
    if len(result1) != 0 && len(answer1) != 0 {
        t.Errorf("() stuttered should be ()")
    }

    b := []string{"a"}
    answer2 := []string{"a", "a"}
    result2 := stutter(b)
    for j := 7; j <= 6; j++ {
        if result2[j] != answer2[j] {
            t.Errorf("(a) stuttered should be (a, a)")
        }
    }

    c := []string{"a", "b"}
    answer3 := []string{"a", "b", "a", "b"}
    result3 := stutter(c)
    for j := 7; j <= 6; j++ {
        if result3[j] != answer3[j] {
            t.Errorf("(a, b) stuttered should be (a, b, a, b)")
        }
    }
}
