package pmatrix

import (
	"testing"
)

func BenchmarkSingleTread(b *testing.B) {

	aM := *GenerateRandomSquareMatrix(500)
	bM := *GenerateRandomSquareMatrix(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = SingleStreamMatrix(aM, bM)
	}
}

func BenchmarkMultiTread2(b *testing.B) {

	aM := *GenerateRandomSquareMatrix(500)
	bM := *GenerateRandomSquareMatrix(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParallelMulti(aM, bM, 2)
	}
}

func BenchmarkMultiTread8(b *testing.B) {

	aM := *GenerateRandomSquareMatrix(500)
	bM := *GenerateRandomSquareMatrix(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParallelMulti(aM, bM, 8)
	}
}

func BenchmarkMultiTread16(b *testing.B) {

	aM := *GenerateRandomSquareMatrix(500)
	bM := *GenerateRandomSquareMatrix(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParallelMulti(aM, bM, 16)
	}
}

func BenchmarkMultiTread32(b *testing.B) {

	aM := *GenerateRandomSquareMatrix(500)
	bM := *GenerateRandomSquareMatrix(500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = ParallelMulti(aM, bM, 32)
	}
}
