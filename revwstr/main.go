package main

import (
	"github.com/01-edu/z01"
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
	tab := SplitWhiteSpaces(os.Args[1])
	var tabrev []string

	for i := len(tab)-1; i >= 0; i-- {
		tabrev = append(tabrev, tab[i], " ")
	}


	for _, v := range tabrev {
		for _, a := range v {
			z01.PrintRune(a)
		}
	}
}