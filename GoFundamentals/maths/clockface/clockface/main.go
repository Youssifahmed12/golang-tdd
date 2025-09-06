package main

import (
	"os"
	"time"

	"github.com/Youssifahmed12/golang-tdd/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
