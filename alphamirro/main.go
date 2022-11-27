package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main()  {
	str := os.Args[1]
	result := ""

	for _, v := range str {
		if v >= 'a' && v <= 'z' {
			v = 'z' - v + 'a'
		} else if v >= 'A' && v <= 'Z' {
			v = 'Z' - v + 'A'
		}
		result += string(v)
	}
	for _, v := range result {
		z01.PrintRune(v)
	}
}
