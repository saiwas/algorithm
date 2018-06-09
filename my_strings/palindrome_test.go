package my_strings

import (
	"testing"
)

func TestPalindrome(t *testing.T)  {
	if !IsPalindrome("madam") {
		t.Error("madam should be true!");
	}

	if IsPalindrome("asdasdqdlqk") {
		t.Error("asdasdqdlqk should be false!");
	}
}
