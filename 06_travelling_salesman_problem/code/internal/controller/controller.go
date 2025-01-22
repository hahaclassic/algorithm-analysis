package controller

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/benchmark"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/graphmap"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp/antcolony"
	"github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/internal/tsp/bruteforce"
	tableoutput "github.com/hahaclassic/algorithm-analysis/06_travelling_salesman_problem/code/pkg/table"
)

type optionHandler struct {
	name string
	f    func() error
}

type Controller struct {
	handlers []*optionHandler
}

func New() *Controller {
	c := &Controller{}
	optionHandlers := []*optionHandler{
		{
			name: "Benchmark",
			f:    c.benchmark,
		},
		{
			name: "Parameterization",
			f:    c.parameterizeAntColony,
		},
		{
			name: "Solve TSP",
			f:    c.solveTSP,
		},
	}

	c.handlers = optionHandlers
	return c
}

func (c *Controller) Run() {
	const exitNumber = 0

	for {
		option := c.getOperationNumber()

		if option == exitNumber {
			break
		}

		if err := c.handlers[option-1].f(); err != nil {
			slog.Error("handler", "err", err)
			continue
		}

		fmt.Println()
	}
}

func (c Controller) getOperationNumber() int {
	c.showMenu()
	fmt.Println()

	var option int
	for {
		fmt.Print("Enter operation number: ")
		if _, err := fmt.Scan(&option); err == nil &&
			option >= 0 && option <= len(c.handlers) {
			break
		}

		fmt.Println("Invalid number. Try again.")
	}

	return option
}

func (c Controller) showMenu() {
	headers := []string{"#", "Operation"}
	rows := make([][]any, 0, len(c.handlers))

	for i := range c.handlers {
		rows = append(rows, []any{i + 1, c.handlers[i].name})
	}
	rows = append(rows, []any{0, "Exit"})

	tableoutput.PrintTable(headers, rows)
}

func (c Controller) readIntWithDefault(prompt string, defaultValue int) int {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil || input == "" {
		return defaultValue
	}
	result, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input, using default value:", defaultValue)
		return defaultValue
	}
	return result
}

func (c Controller) readStringWithDefault(prompt string, defaultValue string) string {
	var input string
	fmt.Print(prompt)
	_, err := fmt.Scanln(&input)
	if err != nil || input == "" {
		return defaultValue
	}
	return input
}

func (c Controller) benchmark() error {
	startCities := c.readIntWithDefault("Enter the start number of cities (default: 3): ", 3)
	endCities := c.readIntWithDefault("Enter the end number of cities (default: 15): ", 15)
	step := c.readIntWithDefault("Enter the step (default: 1): ", 1)
	runs := c.readIntWithDefault("Enter the number of runs (default: 10): ", 10)

	benchmark.BenchmarkTSPSolvers(startCities, endCities, step, runs)

	return nil
}

func (c Controller) parameterizeAntColony() error {
	numIterations := c.readIntWithDefault("Enter the number of iterations (default: 10): ", 10)
	filename := c.readStringWithDefault("Enter the filename (default: ../result/result.txt): ",
		"../result/result.txt")

	params := &antcolony.VariableInputParams{
		Alpha:       []float64{0.1, 0.3, 0.5, 0.7, 0.9},
		Beta:        []float64{0.1, 0.3, 0.5, 0.7, 0.9},
		Evaporation: []float64{0.1, 0.3, 0.5, 0.7, 0.9},
	}

	p := antcolony.NewParameterizer()
	if err := p.ParameterizeAntColony(filename, params, numIterations); err != nil {
		return fmt.Errorf("failed to parameterize ant colony: %v", err)
	}

	slog.Info("Parameterization done.")

	return nil
}

func (c Controller) solveTSP() error {
	filename := c.readStringWithDefault("Enter the filename of the graph (default: ../data/graph_1.json): ",
		"../data/graph_1.json")
	g := &graphmap.Graph{}
	if err := g.LoadFromFile(filename); err != nil {
		return fmt.Errorf("failed to load graph from file: %v", err)
	}

	g.Print()

	bfSolver := &bruteforce.BruteForceSolver{}
	pathBF, timeBF := bfSolver.SolveTSP(g)

	antSolver := antcolony.NewAntColony(nil)
	pathAnt, timeAnt := antSolver.SolveTSP(g)

	printPath := func(path []int) {
		fmt.Print("Path: \n[ ")
		for i := 0; i < len(path)-1; i++ {
			fmt.Print(path[i], " -> ")
		}
		if len(path) >= 1 {
			fmt.Print(path[len(path)-1])
		}
		fmt.Println(" ]")
	}

	fmt.Println("\nBrute Force result:")
	fmt.Println("Best Time:", timeBF)
	printPath(pathBF)

	fmt.Println("\nAnt Colony result:")
	fmt.Println("Best Time:", timeAnt)
	printPath(pathAnt)

	return nil
}
