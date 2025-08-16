package iterations

import "strings"

const repeatTimes = 5

func repeat(s string) string {
	var repeated strings.Builder
	for i := 0; i < repeatTimes; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
