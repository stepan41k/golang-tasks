package main

import (
  "fmt"
)

func longestPalindrome(s string) string {
  if len(s) == 0 {
    return ""
  }

  if len(s) == 1 {
    return s
  }

  res_palindrome := ""

  for i := range s {
    odd_palindrome := is_palindrome(s, i, i + 1)
    even_palindrome := is_palindrome(s, i, i)

    if len(even_palindrome) > len(res_palindrome) {
      res_palindrome = even_palindrome
    }

    if len(odd_palindrome) > len(res_palindrome) {
      res_palindrome = odd_palindrome
    }
  }

  return res_palindrome
}

func is_palindrome(s string, left int, right int) string {
  for left >= 0 && right <= len(s)-1 && s[left] == s[right] {
    left--
    right++
  }

  return s[left+1:right]
}


func main() {
  fmt.Println(longestPalindrome("abacd"))
  fmt.Println(longestPalindrome("dbaabd"))
  fmt.Println(longestPalindrome("db"))
  fmt.Println(longestPalindrome("d"))
  fmt.Println(longestPalindrome(""))
}
