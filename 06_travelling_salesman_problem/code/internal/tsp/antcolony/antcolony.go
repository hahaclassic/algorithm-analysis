package antcolony

import (
	"math"
	"math/rand"
	"time"

	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/graphmap"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

const (
	defaultPheromoneValue float64 = 1.0
)

// var (
// 	waterEvaporationRate  float64 = 0.8
// 	landEvaporationRate   float64 = 0.3
// 	desertEvaporationRate float64 = 0.5
// )

type Settings struct {
	Alpha       float64
	Beta        float64
	Evaporation float64 // Испарение феромонов
	EliteAnts   int
	q           float64
	Iterations  int // Время жизни колонии (tmax)
}

type Colony struct {
	params     Settings
	graph      *graphmap.Graph
	pheromones [][]float64
}

func NewAntColony(settings *Settings) *Colony {
	if settings == nil {
		settings = &Settings{
			Alpha:       0.9,
			Beta:        0.9,
			Evaporation: 0.1,
			EliteAnts:   2,
			q:           100,
			Iterations:  200,
		}
	}

	return &Colony{
		params: *settings,
	}
}

func (c *Colony) NewParams(settings Settings) {
	c.params = settings
}

func (c *Colony) SolveTSP(graph *graphmap.Graph) ([]int, float64) {
	if graph.NumOfCities < 1 {
		return []int{}, 0
	}
	if graph.NumOfCities == 1 {
		return []int{0}, 0
	}

	c.graph = graph
	c.setQ()
	c.initPheromones()
	antPlacement := c.randomAntPlacement()

	bestPath := []int{}
	bestTime := math.MaxFloat64

	for range c.params.Iterations {
		allPaths := make([][]int, 0, c.graph.NumOfCities)
		allTimes := make([]float64, 0, c.graph.NumOfCities)

		for antStartCity := range antPlacement {
			path, travelTime := c.constructPath(antStartCity)

			allPaths = append(allPaths, path)
			allTimes = append(allTimes, travelTime)

			if travelTime < bestTime {
				bestPath = path
				bestTime = travelTime
			}
		}

		c.updatePheromones(allPaths, allTimes, bestPath, bestTime)
	}

	return bestPath, bestTime
}

func (c *Colony) setQ() {
	q := 0.0

	for i := 0; i < c.graph.NumOfCities; i++ {
		for j := 0; j < c.graph.NumOfCities; j++ {
			q += float64(c.graph.TravelTime[i][j])
		}
	}

	q /= float64(c.graph.NumOfCities)
	c.params.q = q
}

func (c *Colony) initPheromones() {
	c.pheromones = make([][]float64, c.graph.NumOfCities)
	for i := range c.pheromones {
		c.pheromones[i] = make([]float64, c.graph.NumOfCities)
		for j := range c.pheromones[i] {
			c.pheromones[i][j] = defaultPheromoneValue
		}
	}
}

func (c *Colony) randomAntPlacement() []int {
	antPlacement := make([]int, c.graph.NumOfCities)
	for i := range antPlacement {
		antPlacement[i] = i
	}

	random.Shuffle(len(antPlacement), func(i, j int) {
		antPlacement[i], antPlacement[j] = antPlacement[j], antPlacement[i]
	})

	return antPlacement
}

func (c *Colony) constructPath(startCity int) ([]int, float64) {
	path := []int{}

	var travelTime float64

	visited := make(map[int]struct{}, c.graph.NumOfCities)
	current := startCity
	visited[current] = struct{}{}
	path = append(path, current)

	for len(visited) < c.graph.NumOfCities {
		next := c.selectNextCity(current, visited)
		if next == -1 {
			break
		}
		visited[next] = struct{}{}
		path = append(path, next)
		travelTime += c.graph.GetTimeOnTerrain(current, next)
		current = next
	}

	return path, travelTime
}

func (c *Colony) selectNextCity(current int, visited map[int]struct{}) int {
	var total float64
	probabilities := make([]float64, c.graph.NumOfCities)

	for next := range c.graph.NumOfCities {
		if _, ok := visited[next]; ok {
			continue
		}
		probabilities[next] = math.Pow(c.pheromones[current][next], c.params.Alpha) *
			math.Pow(1.0/c.graph.GetTimeOnTerrain(current, next), c.params.Beta)
		total += probabilities[next]
	}

	for i := range probabilities {
		probabilities[i] /= total
	}

	randValue := random.Float64()
	accumulated := 0.0
	for next, prob := range probabilities {
		if _, ok := visited[next]; ok {
			continue
		}
		accumulated += prob
		if accumulated >= randValue {
			return next
		}
	}

	return -1
}

func (c *Colony) evaporatePheromones(terrType graphmap.TerrainType, pheromonesValue float64) float64 {
	envInfluence := map[graphmap.TerrainType]float64{
		graphmap.Water:  c.params.Evaporation, //waterEvaporationRate,
		graphmap.Land:   c.params.Evaporation, //landEvaporationRate,
		graphmap.Desert: c.params.Evaporation, //desertEvaporationRate,
	}

	return pheromonesValue * (1 - envInfluence[terrType])
}

func (c *Colony) updatePheromones(allPaths [][]int, allTimes []float64, bestPath []int, bestTime float64) {
	for i := range c.pheromones {
		for j := range c.pheromones[i] {
			c.pheromones[i][j] = c.evaporatePheromones(c.graph.Terrain[i][j], c.pheromones[i][j])
		}
	}

	for i := range allPaths {
		pheromoneDeposit := c.params.q / allTimes[i]
		for j := 0; j < len(allPaths[i])-1; j++ {
			from := allPaths[i][j]
			to := allPaths[i][j+1]

			c.pheromones[from][to] += pheromoneDeposit
			c.pheromones[to][from] += pheromoneDeposit
		}
	}

	elitePheromoneDeposit := float64(c.params.EliteAnts) * c.params.q / bestTime
	for i := 0; i < len(bestPath)-1; i++ {
		from := bestPath[i]
		to := bestPath[i+1]
		c.pheromones[from][to] += elitePheromoneDeposit
		c.pheromones[to][from] += elitePheromoneDeposit
	}
}
