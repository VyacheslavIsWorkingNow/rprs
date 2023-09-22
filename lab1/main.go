package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/rprs/lab1/graphic"
	"github.com/VyacheslavIsWorkingNow/rprs/lab1/pmatrix"
)

const pathBench = `/Users/slavaruswarrior/Documents/GitHub/rprs/lab1/pmatrix/benchmark.txt`

func main() {

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

	graphic.Plot(pathBench)

}
