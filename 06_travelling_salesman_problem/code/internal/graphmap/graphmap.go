package graphmap

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var (
	ErrFileOpen      = errors.New("failed to open the file")
	ErrFileRead      = errors.New("failed to read the file")
	ErrJSONParse     = errors.New("failed to parse JSON")
	ErrInvalidSizes  = errors.New("invalid sizes of travel_time or terrain")
	ErrInvalidMatrix = errors.New("invalid matrix format for travel_time or terrain")
)

type TerrainType int

const (
	Water  TerrainType = iota - 1 // -50% time
	Land                          // +0% time
	Desert                        // +30% time
)

type Graph struct {
	TravelTime  [][]float64     `json:"travel_time"`
	Terrain     [][]TerrainType `json:"terrain"`
	NumOfCities int             `json:"num_of_cities"`
}

func (g Graph) GetTimeOnTerrain(from, to int) float64 {
	envInfluence := map[TerrainType]float64{
		Water:  0.5,
		Land:   1,
		Desert: 1.3,
	}

	return g.TravelTime[from][to] * envInfluence[g.Terrain[from][to]]
}

func (g *Graph) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFileOpen, err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrFileRead, err)
	}

	if err := json.Unmarshal(data, g); err != nil {
		return fmt.Errorf("%w: %s", ErrJSONParse, err)
	}

	if len(g.TravelTime) != g.NumOfCities || len(g.Terrain) != g.NumOfCities {
		return ErrInvalidSizes
	}
	for i := 0; i < g.NumOfCities; i++ {
		if len(g.TravelTime[i]) != g.NumOfCities || len(g.Terrain[i]) != g.NumOfCities {
			return ErrInvalidMatrix
		}
	}

	return nil
}

func (g *Graph) GenerateRandom(numOfCities int, minTime, maxTime float64) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	g.TravelTime = make([][]float64, numOfCities)
	g.Terrain = make([][]TerrainType, numOfCities)
	g.NumOfCities = numOfCities

	for i := 0; i < numOfCities; i++ {
		g.TravelTime[i] = make([]float64, numOfCities)
		g.Terrain[i] = make([]TerrainType, numOfCities)
		for j := 0; j < numOfCities; j++ {
			if i != j {
				g.TravelTime[i][j] = float64(random.Intn(10) + 1)
				terrainType := TerrainType(random.Intn(3) - 1)
				g.Terrain[i][j] = terrainType
			} else {
				g.TravelTime[i][j] = 0
				g.Terrain[i][j] = Land
			}
		}
	}
}

func (g *Graph) Print() {
	fmt.Println("\nNumber of cities:", g.NumOfCities)

	fmt.Println("\nTravel Time Matrix:")
	for i, row := range g.TravelTime {
		fmt.Printf("City %d: ", i)
		for _, time := range row {
			fmt.Printf("%6.2f ", time)
		}
		fmt.Println()
	}

	fmt.Println("\nTerrain Matrix:")
	for i, row := range g.Terrain {
		fmt.Printf("City %d: ", i)
		for _, terrain := range row {
			var terrainName string
			switch terrain {
			case Water:
				terrainName = "w"
			case Land:
				terrainName = "l"
			case Desert:
				terrainName = "d"
			default:
				terrainName = "Unknown"
			}
			fmt.Printf("%-2s ", terrainName)
		}
		fmt.Println()
	}
}
