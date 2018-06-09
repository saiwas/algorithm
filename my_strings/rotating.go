package my_strings;

import "fmt"
import "time"

func Rotating() {
	start_time := time.Now();
	result := RotatingString1("abcdef", 3);
	fmt.Println(result);
	end_time := time.Now();
	fmt.Println(end_time.Sub(start_time));  // 64.255µs

	start_time2 := time.Now();
	result2 := RotatingString2("abcdef", 3);
	fmt.Println(result2);
	end_time2 := time.Now();
	fmt.Println(end_time2.Sub(start_time2));    // 3.41µs
}

// 强拉
func RotatingString1(s string, m int) string {
	bytes := []byte(s);

	for j := 0; j < m; j++ {
		tmp := bytes[0];
		n := len(bytes);
		for i := 1; i < n; i++ {
			bytes[i - 1] = bytes[i];
		}

		bytes[n - 1] = tmp;
	}
	return string(bytes);
}

func ReverseString(bs []byte, from int, to int)  {
	var tmp byte;

	for from < to {
		tmp = bs[from];
		bs[from] = bs[to];
		bs[to] = tmp;
		from++;
		to--;
	}
}

func RotatingString2(s string, m int) string {
	n := len(s);
	bs := []byte(s);

	ReverseString(bs, 0, m - 1);
	ReverseString(bs, m, n - 1);
  ReverseString(bs, 0, n - 1);

	return string(bs);
}

