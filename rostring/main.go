package main

import (
	"fmt"
	"os"
)

func SplitWhiteSpaces(s string) []string  {
	var tab []string
	mot := ""

	for i, v := range s {
		if v > 32 {
			mot += string(v)
		} else if mot != "" {
			tab = append(tab, mot)
			mot = ""
		}
		if len(s)-1 == i {
			tab = append(tab, mot)
		}
	}
	return tab
}

func main() {
	
	size := len(os.Args)

	if size == 2 {
		str := SplitWhiteSpaces(os.Args[1])

		for _, v := range str[1:] {
			fmt.Printf(v) 
			fmt.Printf(" ")
		}
		fmt.Printf(str[0])
	}

	
}