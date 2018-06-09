package my_strings;

import (
	"bytes"
	"strings"
)

func Manacher(s string) int {
	clean_up_s := clean_up_string(s);
	s_length := len(clean_up_s);
	p := make([]int, s_length);  // 记录对称长度的数组
	t := []byte(clean_up_s);
	var c = 0;   // 标记每个回文的对称中心
	var r = 0;   // 标记回文subtring的最右端
	var mirr int;

	for i := 1; i < s_length - 1; i++ {
		mirr = 2 * c - i;

		if i < r {
			p[i] = Min(r - i, p[mirr]);
		}

		for t[i + (1 + p[i])] == t[i - (1 + p[i])] {
			p[i]++;
		}

		if i + p[i] > r {
      c = i;
      r = i + p[i];
    }
	}

	return Max(p);
}

func Max(array []int) int {
	max := array[0];

	for i := 1; i < len(array); i++ {
		if max < array[i] {
			max = array[i];
		}
	}
	return max;
}

func Min(a int, b int) int {
	if a < b {
		return a;
	}
	return b;
}

func clean_up_string(s string) string {
	var buffer bytes.Buffer;
	buffer.WriteString("$#");

	s_array := strings.Split(s, "");

	for i := 0; i < len(s_array); i++ {
		buffer.WriteString(s_array[i]);
		buffer.WriteString("#");
	}

	buffer.WriteString("@");

	return buffer.String();
}
