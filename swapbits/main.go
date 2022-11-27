package main

import "fmt"

func main()  {
	fmt.Println(SwapBits(0100))
}

func SwapBits(octet byte) byte {
	return octet<<4 | octet>>4
}