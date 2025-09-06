package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, s string) {
	fmt.Fprintf(writer, "Hello, %s", s)
}
