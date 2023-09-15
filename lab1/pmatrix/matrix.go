package pmatrix

import (
	"fmt"
	"math"
	"sync"
)

type Matrix struct {
	buffer     [][]int
	rows, cols int // rows - строки, cols - столбцы
}

func InitMatrix(r, c int) *Matrix {
	matrix := Matrix{
		buffer: make([][]int, r),
		rows:   r,
		cols:   c,
	}

	for i := 0; i < matrix.rows; i++ {
		matrix.buffer[i] = make([]int, matrix.cols)
	}

	return &matrix
}

func (m *Matrix) AddBuffer(buf []int) error {

	if len(buf) != m.cols*m.rows {
		return fmt.Errorf("failed add buffer: sized missed")
	}

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			m.buffer[i][j] = buf[i*m.rows+j]
		}
	}

	return nil
}

func SingleStreamMatrix(a, b Matrix) (Matrix, error) {
	if a.cols != b.rows {
		return Matrix{}, fmt.Errorf("can't multi matrix, where a.cols != b.rows")
	}

	ans := *InitMatrix(a.rows, b.cols)

	calculateChunk(ans, a, b, 0, a.rows)

	return ans, nil
}

func (m *Matrix) String() string {
	ans := "[\n"

	for i := 0; i < m.rows; i++ {
		ans += fmt.Sprintln(m.buffer[i])
	}

	ans += "]"

	return ans
}

func ParallelMulti(a, b Matrix, chunks int) (Matrix, error) {
	if a.cols != a.rows || b.cols != b.rows {
		return Matrix{}, fmt.Errorf("can't multiply non-square matrices")
	}
	if a.cols != b.cols {
		return Matrix{}, fmt.Errorf("can't multiply matrices with unequal dimensions")
	}
	if a.cols%(int(math.Sqrt(float64(chunks)))) != 0 {
		return Matrix{}, fmt.Errorf("can't multiply matrices with the given number of chunks")
	}

	wg := &sync.WaitGroup{}

	result := *InitMatrix(a.rows, b.cols)

	rowsPerChunk := a.rows / chunks

	for i := 0; i < chunks; i++ {
		startRow := i * rowsPerChunk
		endRow := (i + 1) * rowsPerChunk

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			calculateChunk(result, a, b, start, end)
		}(startRow, endRow)
	}

	wg.Wait()

	return result, nil
}

func calculateChunk(result, a, b Matrix, start, end int) {
	for i := start; i < end; i++ {
		for j := 0; j < b.cols; j++ {
			for k := 0; k < a.cols; k++ {
				result.buffer[i][j] += a.buffer[i][k] * b.buffer[k][j]
			}
		}
	}
}
