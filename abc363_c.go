abc363_c.go
#################################################
[TLE]

package main

import (
    "fmt"
)

// Check if any substring of length k is a palindrome
func isNonPalindromic(s string, k int) bool {
    for i := 0; i <= len(s)-k; i++ {
        substring := s[i : i+k]
        if isPalindrome(substring) {
            return false
        }
    }
    return true
}

// Check if a string is a palindrome
func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

// Generate all unique permutations of a string
func permute(str string) []string {
    var result []string
    used := make(map[string]bool)

    var helper func([]rune, int)
    helper = func(s []rune, idx int) {
        if idx == len(s) {
            perm := string(s)
            if !used[perm] {
                result = append(result, perm)
                used[perm] = true
            }
            return
        }
        for i := idx; i < len(s); i++ {
            s[idx], s[i] = s[i], s[idx]
            helper(s, idx+1)
            s[idx], s[i] = s[i], s[idx]
        }
    }

    helper([]rune(str), 0)
    return result
}

func main() {
    var N, K int
    var S string
    fmt.Scan(&N, &K)
    fmt.Scan(&S)

    perms := permute(S)
    count := 0

    for _, perm := range perms {
        if isNonPalindromic(perm, K) {
            count++
        }
    }

    fmt.Println(count)
}
