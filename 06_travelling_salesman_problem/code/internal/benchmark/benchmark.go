package benchmark

import (
	"time"

	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/graphmap"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp/antcolony"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp/bruteforce"
	tableoutput "github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/pkg/table"
)

func BenchmarkTSPSolvers(startCities int, endCities int, step int, runs int) {
	bfSolver := &bruteforce.BruteForceSolver{}
	colony := antcolony.NewAntColony(nil)

	headers := []string{"num of cities", "brute force", "ant colony"}
	rows := [][]interface{}{}

	for numCities := startCities; numCities <= endCities; numCities += step {
		bfDuration := benchmark(bfSolver, numCities, runs)
		colonyDuration := benchmark(colony, numCities, runs)

		rows = append(rows, []interface{}{numCities, bfDuration, colonyDuration})
	}

	tableoutput.PrintTable(headers, rows)
}

func benchmark(solver tsp.TSPSolver, cities int, runs int) time.Duration {
	var elapsed time.Duration

	for range runs {
		graph := &graphmap.Graph{}
		graph.GenerateRandom(cities, 10, 100)
		startTime := time.Now()
		solver.SolveTSP(graph)
		elapsed += time.Since(startTime)
	}

	return elapsed / time.Duration(runs)
}
