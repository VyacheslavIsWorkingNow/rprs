all: run


build:
	@go build -tags mpi ./parallel/main.go

# example: make run p=2 n=1000
run: build
	@mpirun -np $(p) ./main -n=$(n)

clear:
	@rm ./main

