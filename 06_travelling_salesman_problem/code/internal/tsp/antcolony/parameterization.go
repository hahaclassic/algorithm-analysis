package antcolony

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/graphmap"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp/bruteforce"
)

var (
	ErrFailedParameterization = errors.New("failed parameterization")
)

const (
	graphsDir = "../data"
)

type VariableInputParams struct {
	Alpha       []float64
	Beta        []float64
	Evaporation []float64
}

type ParamResult struct {
	Alpha       float64
	Beta        float64
	Evaporation float64
	Deviation   *Deviation
}

type Deviation struct {
	MaxDeviation    float64
	MedianDeviation float64
	AvgDeviation    float64
}

type Parameterizer struct {
	Data        []*ParamResult
	RunsCount   int
	GraphsCount int
}

func NewParameterizer() *Parameterizer {
	return &Parameterizer{}
}

func (p *Parameterizer) ParameterizeAntColony(filename string, params *VariableInputParams, numRuns int) error {
	graphs, err := p.loadGraphs()
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedParameterization, err)
	}
	p.GraphsCount = len(graphs)
	p.RunsCount = numRuns

	bruteforcer := &bruteforce.BruteForceSolver{}
	bestTimes := make([]float64, p.GraphsCount)
	for i := range p.GraphsCount {
		_, travelTime := bruteforcer.SolveTSP(graphs[i])
		bestTimes[i] = travelTime
	}

	for _, alpha := range params.Alpha {
		for _, beta := range params.Beta {
			for _, evaporation := range params.Evaporation {
				colony := NewAntColony(&Settings{
					Alpha:       alpha,
					Beta:        beta,
					Evaporation: evaporation,
					EliteAnts:   10,
					Q:           100,
					Iterations:  10,
				})
				deviation := p.calc(colony, graphs, bestTimes)

				p.Data = append(p.Data, &ParamResult{
					Alpha:       alpha,
					Beta:        beta,
					Evaporation: evaporation,
					Deviation:   deviation,
				})
			}
		}
	}

	err = p.saveToFile(filename)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrFailedParameterization, err)
	}

	return nil
}

func (p *Parameterizer) calc(colony *Colony, graphs []*graphmap.Graph, bestTimes []float64) *Deviation {
	deviations := make([]float64, 0, len(graphs)*p.RunsCount)

	for i := range graphs {
		for range p.RunsCount {
			_, travelTime := colony.SolveTSP(graphs[i])
			deviations = append(deviations, travelTime-bestTimes[i])
		}
	}

	sort.Float64s(deviations)

	return &Deviation{
		MaxDeviation:    deviations[len(deviations)-1],
		MedianDeviation: deviations[len(deviations)/2],
		AvgDeviation:    p.calculateAvgDeviation(deviations),
	}
}

func (Parameterizer) loadGraphs() ([]*graphmap.Graph, error) {
	graphs := []*graphmap.Graph{}

	err := filepath.Walk(graphsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil
		}

		graph := &graphmap.Graph{}
		err = graph.LoadFromFile(path)
		graphs = append(graphs, graph)

		return err
	})

	if err != nil {
		return nil, err
	}

	return graphs, nil
}

func (Parameterizer) calculateAvgDeviation(deviations []float64) float64 {
	var sum float64
	for _, d := range deviations {
		sum += d
	}

	return float64(sum) / float64(len(deviations))
}

func (p *Parameterizer) saveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString("Alpha  Beta  Evaporation  MaxDeviation  MedianDeviation  AvgDeviation\n")
	if err != nil {
		return fmt.Errorf("failed to write header to file: %v", err)
	}

	for _, param := range p.Data {
		line := fmt.Sprintf("%.2f  %.2f  %.2f  %.2f  %.2f  %.2f\n",
			param.Alpha,
			param.Beta,
			param.Evaporation,
			param.Deviation.MaxDeviation,
			param.Deviation.MedianDeviation,
			param.Deviation.AvgDeviation,
		)
		_, err = file.WriteString(line)
		if err != nil {
			return fmt.Errorf("failed to write data to file: %v", err)
		}
	}

	_, err = file.WriteString(fmt.Sprintf("\nRuns Count: %d  Graphs Count: %d\n", p.RunsCount, p.GraphsCount))
	if err != nil {
		return fmt.Errorf("failed to write counts to file: %v", err)
	}

	return nil
}
