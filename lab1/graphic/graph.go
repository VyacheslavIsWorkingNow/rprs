package graphic

import (
	"bufio"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"os"
	"regexp"
)

func Plot(path string) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	regex := regexp.MustCompile(`Benchmark(\w+)-\d+\s+(\d+)\s+(\d+)\s+ns/op`)

	names := make([]string, 0)
	ops := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := regex.FindStringSubmatch(line)
		if len(match) == 4 {
			benchmarkName := match[1]
			opsPerNanoSecond := match[3]

			var opsValue int
			_, _ = fmt.Sscanf(opsPerNanoSecond, "%d", &opsValue)
			names = append(names, benchmarkName)
			ops = append(ops, opsValue)
		}
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	p := plot.New()

	points := make(plotter.Values, len(ops))
	for i, v := range ops {
		points[i] = float64(v)
	}

	barchart, err := plotter.NewBarChart(points, vg.Points(40))
	if err != nil {
		panic(err)
	}

	p.NominalX(names...)

	p.Add(barchart)

	if err = p.Save(7*vg.Inch, 7*vg.Inch, "benchmark.png"); err != nil {
		panic(err)
	}
}
