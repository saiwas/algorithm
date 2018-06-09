package my_strings;

import (
	"fmt"
	"strings"
)

func Palindrome()  {
	fmt.Printf("PalindromeTest");
}

func IsPalindrome(s string) bool {
	lower_s := strings.ToLower(s);
	front := len(lower_s) / 2;
	back := front;

	for front > 0 {
		if lower_s[front] != lower_s[back] {
      return false;
    }

    front--;
    back++;
	}

	return true;
}
