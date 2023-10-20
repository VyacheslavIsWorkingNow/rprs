package main

import (
	"context"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/rprs/lab2/iterations"
	"gonum.org/v1/gonum/mat"
	"time"
)

func main() {
	N := 4

	A := mat.NewDense(N, N, nil)
	for i := 0; i < N; i++ {
		A.Set(i, i, 2.0)
		for j := 0; j < N; j++ {
			if i != j {
				A.Set(i, j, 1.0)
			}
		}
	}

	bData := make([]float64, N)
	for i := 0; i < N; i++ {
		bData[i] = float64(N + 1)
	}
	b := mat.NewVecDense(N, bData)

	xData := make([]float64, N)
	x := mat.NewVecDense(N, xData)

	epsilon := 0.00001

	fmt.Println("Матрица A:")
	fmt.Println(mat.Formatted(A))
	fmt.Println("Вектор b:", b)
	fmt.Println("Начальное приближение x:", x)
	fmt.Println("Желаемая точность:", epsilon)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result := iterations.SimpleIteration(ctx, A, b, x, epsilon)

	fmt.Println("Результат:")
	fmt.Printf("%+v", mat.Formatted(result))
}
