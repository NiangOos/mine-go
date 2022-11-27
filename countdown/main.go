package main

import "github.com/01-edu/z01"

func main()  {
	str := "9876543210"

	for _, v := range str {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}