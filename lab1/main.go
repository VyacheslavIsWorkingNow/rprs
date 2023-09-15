package main

import (
	"fmt"
	"lab1/pmatrix"
)

func main() {
	fmt.Println("Hi")

	// На потоки делим только тогда, когда они делятся нацело

	a := *pmatrix.InitMatrix(4, 4)
	b := *pmatrix.InitMatrix(4, 4)

	aErr := a.AddBuffer([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	})
	if aErr != nil {
		panic("a missed")
	}

	bErr := b.AddBuffer([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
	})
	if bErr != nil {
		panic("b missed")
	}

	c, _ := pmatrix.SingleStreamMatrix(a, b)

	fmt.Println(c.String())

	pc, _ := pmatrix.ParallelMulti(a, b, 2)

	fmt.Println(pc.String())

}

/* [
[30 36 42]
[66 81 96]
[102 126 150]
]

[
[90 100 110 120]
[202 228 254 280]
[314 356 398 440]
[426 484 542 600]
]

*/
