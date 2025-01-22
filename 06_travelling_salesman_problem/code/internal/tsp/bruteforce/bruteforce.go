package bruteforce

import (
	"math"

	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/graphmap"
)

type BruteForceSolver struct {
	graph *graphmap.Graph
}

func (b *BruteForceSolver) SolveTSP(graph *graphmap.Graph) ([]int, float64) {
	if graph.NumOfCities < 1 {
		return []int{}, 0
	}
	if graph.NumOfCities == 1 {
		return []int{0}, 0
	}
	b.graph = graph

	cities := make([]int, graph.NumOfCities)
	for i := range cities {
		cities[i] = i
	}

	bestPath := []int{}
	bestTime := math.MaxFloat64

	permute(cities, func(path []int) {
		travelTime := b.calculateTravelTime(path)

		if travelTime < bestTime {
			bestTime = travelTime
			bestPath = make([]int, len(path))
			copy(bestPath, path)
		}
	})

	return bestPath, bestTime
}

func (b *BruteForceSolver) calculateTravelTime(route []int) float64 {
	var distance float64
	for i := 0; i < len(route)-1; i++ {
		distance += b.graph.GetTimeOnTerrain(route[i], route[i+1])
	}

	return distance
}

func permute(arr []int, callback func([]int)) {
	var generate func([]int, int)
	generate = func(a []int, n int) {
		if n == 1 {
			callback(a)
			return
		}
		for i := 0; i < n; i++ {
			generate(a, n-1)
			if n%2 == 1 {
				a[0], a[n-1] = a[n-1], a[0]
			} else {
				a[i], a[n-1] = a[n-1], a[i]
			}
		}
	}

	generate(arr, len(arr))
}
