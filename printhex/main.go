package main

import (
	"github.com/01-edu/z01"
	"strconv"
	"os"
)

func main() {
	ar := os.Args

	if len(ar) == 2 {
		n, er := strconv.Atoi(ar[1])
		if er != nil {
			z01.PrintRune('0')
		}

		for _, r := range hex(n) {
			z01.PrintRune(r)
		}

	}
	z01.PrintRune('\n')
}

func hex(n int) string {
	base := "0123456789abcdef"
	if n < 16 {
		return string(base[n])
	}
	return hex(n/16) + string(base[n%16])
}