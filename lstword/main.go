package main

import (
	"os"
	"github.com/01-edu/z01"
)

func SplitWithSpace(s string) []string  {
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

func main()  {
	size := len(os.Args) 
	
	if size > 1 {
		tab := SplitWithSpace(os.Args[1])
		last := tab[len(tab)-1]
		for _, v := range last {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}

}
