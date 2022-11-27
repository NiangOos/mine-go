package main

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	var tab []string
	mot := ""

	for i, str := range s {
		if str > 32 {
			mot += string(str)
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
	tab := SplitWhiteSpaces("Hello how are you?")
	size := len(tab)
	var new []string
	for i := size - 1; i >= 0; i-- {
		new = append(new, tab[i])
	}

	for _, v := range tab {
		fmt.Printf(v)
	}

	
}


// split splits


package piscine

func Split(str, charset string) []string {
	resp := []string{}
	new := ""
	for i := 0; i < len(str); i++ {
		if Charset(str, charset, i) {
			if new != "" {
				resp = append(resp, new)
				new = ""
				i = i + len(charset) - 1
			}
		} else {
			new = new + string(str[i])
		}
	}
	if new != "" {
		resp = append(resp, new)
	}
	return resp
}

func Charset(str, charset string, i int) bool {
	j := 0
	for j < len(charset) && i < len(str) {
		if str[i] != charset[j] {
			return false
		}
		i++
		j++
	}
	return true
}

