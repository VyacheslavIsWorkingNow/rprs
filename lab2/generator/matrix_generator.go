package generator

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func GenMatrix(N int) *mat.Dense {
	A := mat.NewDense(N, N, nil)
	for i := 0; i < N; i++ {
		A.Set(i, i, 2.0)
		for j := 0; j < N; j++ {
			if i != j {
				A.Set(i, j, 1.0)
			}
		}
	}
	return A
}

func GenVectorConstN(N int) []float64 {
	bData := make([]float64, N)
	for i := 0; i < N; i++ {
		bData[i] = float64(N + 1)
	}
	return bData
}

func GenChunkMatrix(start, end, size int) *mat.Dense {
	rows := end - start
	res := mat.NewDense(rows, size, make([]float64, rows*size))

	for i := start; i < end; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				res.Set(i-start, j, 2.0)
			} else {
				res.Set(i-start, j, 1.0)
			}
		}
	}

	return res
}

func GenCheckFreeVector(matrixChunkRowsSize, N int) *mat.VecDense {
	resBuf := make([]float64, matrixChunkRowsSize)
	for i := 0; i < matrixChunkRowsSize; i++ {
		resBuf[i] = float64(N + 1)
	}
	return mat.NewVecDense(matrixChunkRowsSize, resBuf)
}

func GenRandomVector(N int) []float64 {
	res := make([]float64, N)

	for i := 0; i < N; i++ {
		res[i] = float64(rand.Intn(N))
	}

	return res
}
