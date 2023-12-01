package main

import (
	"context"
	"fmt"
	"time"

	"gonum.org/v1/gonum/mat"

	"github.com/VyacheslavIsWorkingNow/rprs/lab2/generator"
	"github.com/VyacheslavIsWorkingNow/rprs/lab2/iterations"
)

var size = 10000

func main() {

	startTime := time.Now()
	A := generator.GenMatrix(size)
	bData := generator.GenVectorConstN(size)

	b := mat.NewVecDense(size, bData)

	xData := generator.GenRandomVector(size)
	x := mat.NewVecDense(size, xData)

	epsilon := 1e-5

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	_ = iterations.SimpleIteration(ctx, A, b, x, epsilon)
	// printResult(A, result, b, x, epsilon)

	fmt.Println("Program time:", time.Since(startTime))
}

func printResult(A *mat.Dense, result, b, x *mat.VecDense, epsilon float64) {
	fmt.Println("Матрица A:")
	fmt.Println(mat.Formatted(A))
	fmt.Println("Вектор b:", b)
	fmt.Println("Начальное приближение x:", x)
	fmt.Println("Желаемая точность:", epsilon)
	fmt.Println("Результат:")
	fmt.Printf("%+v\n", mat.Formatted(result))
}
