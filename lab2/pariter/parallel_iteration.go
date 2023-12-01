package pariter

import (
	"context"
	"log"
	"math"

	"github.com/emer/empi/mpi"
	"gonum.org/v1/gonum/mat"
)

type mpiObj struct {
	comm  *mpi.Comm
	rank  int
	start int
	end   int
}

type matrixChunk struct {
	baseMatrixChunk *mat.Dense
	freeConst       *mat.VecDense
	size            int
}

type SolverWithVecSeparation struct {
	mo  *mpiObj
	mc  *matrixChunk
	tau float64
	eps float64
}

func NewSolverWithVecSeparation(
	comm *mpi.Comm, baseMatrixChunk *mat.Dense, freeConst *mat.VecDense, size int, eps float64,
) *SolverWithVecSeparation {
	mo := newMpiObj(comm, size)
	t := float64(1) / float64(size)
	mc := newMatrixChunk(baseMatrixChunk, freeConst, size)
	return &SolverWithVecSeparation{
		mo:  mo,
		mc:  mc,
		eps: eps,
		tau: t,
	}
}

func newMpiObj(comm *mpi.Comm, size int) *mpiObj {
	rowStart, rowEnd := GetMpiChunkParams(comm, size)
	rank := comm.Rank()
	return &mpiObj{
		rank:  rank,
		start: rowStart,
		end:   rowEnd,
		comm:  comm,
	}
}

func GetMpiChunkParams(comm *mpi.Comm, size int) (start, end int) {
	chunkSize := size / comm.Size()
	start = comm.Rank() * chunkSize
	end = (comm.Rank() + 1) * chunkSize
	return
}

func newMatrixChunk(baseMatrixChunk *mat.Dense, freeConst *mat.VecDense, size int) *matrixChunk {
	return &matrixChunk{
		baseMatrixChunk: baseMatrixChunk,
		freeConst:       freeConst,
		size:            size,
	}
}

func (s *SolverWithVecSeparation) FindSolution(ctx context.Context) *mat.VecDense {

	res := mat.NewVecDense(s.mo.end-s.mo.start, nil)
	buff := make([]float64, s.mc.size)
	iteration := 0

	s.iterateLoop(ctx, res, &buff, &iteration)

	return res
}

func (s *SolverWithVecSeparation) iterateLoop(ctx context.Context, res *mat.VecDense, buff *[]float64, iteration *int) {
	metric := math.MaxFloat64
	prev := metric
	for metric > s.eps {
		select {
		case <-ctx.Done():
			log.Printf("Context deadline\n")
			return
		default:
		}
		*iteration++
		chunk := s.iterateSolution(res).RawVector().Data

		_ = s.mo.comm.AllGatherF64(*buff, chunk)
		tmp := mat.NewVecDense(s.mc.size, *buff)

		res.SubVec(res, tmp.SliceVec(s.mo.start, s.mo.end))

		prev = metric
		metric = s.calcMetric(res)

		if prev < metric {
			s.tau *= -1
		}
	}
}

func (s *SolverWithVecSeparation) iterateSolution(chunk *mat.VecDense) *mat.VecDense {
	buff := make([]float64, s.mc.size)
	_ = s.mo.comm.AllGatherF64(buff, chunk.RawVector().Data)

	tmp := mat.NewVecDense(s.mo.end-s.mo.start, nil)

	tmp.MulVec(s.mc.baseMatrixChunk, mat.NewVecDense(s.mc.size, buff))
	tmp.SubVec(tmp, s.mc.freeConst)
	tmp.ScaleVec(s.tau, tmp)

	return tmp
}

type fraction struct {
	numerator   float64
	denominator float64
}

func (s *SolverWithVecSeparation) calcMetric(chunk *mat.VecDense) float64 {
	buff := make([]float64, s.mc.size)
	_ = s.mo.comm.AllGatherF64(buff, chunk.RawVector().Data)

	tmp := mat.NewVecDense(s.mo.end-s.mo.start, nil)

	tmp.MulVec(s.mc.baseMatrixChunk, mat.NewVecDense(s.mc.size, buff))
	tmp.SubVec(tmp, s.mc.freeConst)

	frac := getChunkMetricFraction(s, tmp)

	sumFrac := make([]float64, 2)
	_ = s.mo.comm.AllReduceF64(mpi.OpSum, sumFrac, []float64{frac.numerator, frac.denominator})
	return getMetricValue(sumFrac)
}

func getChunkMetricFraction(s *SolverWithVecSeparation, tmp *mat.VecDense) fraction {
	f := fraction{}

	for _, v := range tmp.RawVector().Data {
		f.numerator += v * v
	}
	for _, v := range s.mc.freeConst.RawVector().Data {
		f.denominator += v * v
	}
	return f
}

func getMetricValue(frac []float64) float64 {
	return math.Sqrt(frac[0]) / math.Sqrt(frac[1])
}
