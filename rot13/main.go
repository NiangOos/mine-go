package main

import (
	"os"
	"github.com/01-edu/z01"
)

func Rot13(s string) string{
	result := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if v > 'm' {
				v -= 13
			} else {
				v += 13
			}
		} else if v >= 'A' && v <= 'Z' {
			if v > 'M' {
				v -= 13
			} else {
				v += 13
			}
		}
		result += string(v)
	} 
	return result
}

func main()  {
	output := Rot13(os.Args[1])
	for _, v := range output {
		z01.PrintRune(v)
	}
}