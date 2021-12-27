package main

import "fmt"

func main() {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = 100 + i
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/*
Element[0] = 100
Element[1] = 101
Element[2] = 102
Element[3] = 103
Element[4] = 104
Element[5] = 105
Element[6] = 106
Element[7] = 107
Element[8] = 108
Element[9] = 109
*/
