package main

import (
	"os"

	"github.com/01-edu/z01"
)

func match(str1 string, str2 string) bool {
	var match string
	count := 0
	for i := 0; i < len(str1); i++ {
		for j := count; j < len(str2); j++ {
			if str1[i] == str2[j] {
				match += string(str1[i])
				count = j + 1
				break
			}
		}
	}
	if match != str1 {
		return false 
	} 
	return true

}

func main() {

if len(os.Args) == 3 {
	if match(os.Args[1], os.Args[2]) {
		for _, v := range os.Args[1] {
			z01.PrintRune(v)
		} 
		z01.PrintRune('\n')
	}
}
	
}

