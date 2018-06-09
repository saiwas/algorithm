package my_strings

import (
	"testing"
)

func TestRotating(t *testing.T)  {
	except_result := "defabc";

	test_result := RotatingString1("abcdef", 3);
	if test_result != except_result {
		t.Errorf("RotatingString1 got error, except_result: %s, but we got: %s", except_result, test_result);
	}

	test_result = RotatingString2("abcdef", 3);
	if test_result != except_result {
		t.Errorf("RotatingString1 got error, except_result: %s, but we got: %s", except_result, test_result);
	}
}
