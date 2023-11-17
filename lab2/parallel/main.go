package main

import (
	"flag"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/rprs/lab2/generator"
	"github.com/VyacheslavIsWorkingNow/rprs/lab2/pariter"
	"github.com/emer/empi/mpi"
	"log"
	"time"
)

const (
	eps = 1e-5
)

var (
	size = 10
)

func init() {
	flag.IntVar(&size, "n", size, "size of matrix")
}

func main() {

	startTime := time.Now()
	flag.Parse()
	defer mpi.Finalize()

	mpi.Init()
	comm, errC := mpi.NewComm(nil)
	if errC != nil {
		log.Fatalf("can't init mpi comm %e", errC)
	}

	start, end, buff := initializationMpiMatrix(comm)

	errB := comm.BcastF64(0, buff)
	if errB != nil {
		log.Fatalf("can't broadcast buf %e", errB)
	}

	matrixChunk := generator.GenChunkMatrix(start, end, size)
	freeConst := generator.GenCheckFreeVector(matrixChunk.RawMatrix().Rows, size)
	solver := pariter.NewSolverWithVecSeparation(comm, matrixChunk, freeConst, size, eps)

	s := solver.FindSolution()
	fullRes := make([]float64, size)

	errA := comm.AllGatherF64(fullRes, s.RawVector().Data)
	if errA != nil {
		log.Fatalf("can't record res data %e", errA)
	}

	defer fmt.Println("Success")
	fmt.Println("Program time:", time.Since(startTime))
}

func initializationMpiMatrix(comm *mpi.Comm) (start int, end int, buff []float64) {

	chunkSize := size / comm.Size()
	start = comm.Rank() * chunkSize
	end = (comm.Rank() + 1) * chunkSize

	buff = make([]float64, size)

	if isStreamZero(comm.Rank()) {
		buff = generator.GenRandomVector(size)
	}

	return
}

func isStreamZero(stream int) bool {
	return stream == 0
}