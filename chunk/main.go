package main

import "fmt"

func Chunk(slice []int, Size int) {

	if Size == 0 {
		fmt.Println()
	} else {
		var chunks [][]int

		for i := 0; i < len(slice); i += Size {
			end := i + Size

			if end > len(slice) {
				end = len(slice)
			}

			chunks = append(chunks, slice[i:end])
		}
		fmt.Println(chunks)
	}

}

func main() {
	//Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}