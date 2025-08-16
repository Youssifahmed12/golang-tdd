package iterations

import "strings"

func Repeat(s string, n int) string {
	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
