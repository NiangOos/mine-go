package main

import "fmt"

func FoldInt(f func(int, int) int, a []int, n int) {
	ac := n
	for _, v := range a {
		ac = f(ac, v)
	}
	fmt.Println(ac)
}

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
