package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	size := len(os.Args)

	if size > 1 {
		last := os.Args[size-1]

		for _, v := range last {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
		
	}
}