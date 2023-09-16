all:

.PHONY: run
run:
	go run .

.PHONY: benchmarkFast
benchmarkFast:
	cd pmatrix && go test -bench=. -benchmem -benchtime=10x

.PHONY: benchmarkSlow
benchmarkSlow:
	cd pmatrix && go test -bench=. -benchmem -benchtime=100x

.PHONY: saveBenchResult
saveBenchResult:
	cd pmatrix && go test -bench=. > benchmark.txt



