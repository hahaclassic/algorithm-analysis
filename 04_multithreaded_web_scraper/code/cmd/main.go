package main

import (
	"flag"
	"fmt"

	scraper "github.com/hahaclassic/algorithm-analysis/04_multithreaded_web_scraper/code/internal/scraper"
)

var (
	baseURL    string
	dirPath    string
	maxWorkers int
	maxPages   int
	bench      bool
	repeats    int
)

func init() {
	flag.BoolVar(&bench, "bench", false, "turn on benchmark")
	flag.IntVar(&repeats, "repeats", 1, "test repeats")
	flag.StringVar(&baseURL, "url", "https://edimdoma.ru", "the main page of the web resource")
	flag.StringVar(&dirPath, "dir", "../data", "the directory where the received html pages are saved")
	flag.IntVar(&maxWorkers, "workers", 1, "max number of goroutines")
	flag.IntVar(&maxPages, "pages", 10, "max number of pages")
	flag.Parse()
}

func main() {
	var visited int

	fmt.Println("[PARAMETERS]",
		"\n- baseURL:", baseURL,
		"\n- dirPath:", dirPath,
		"\n- maxPages:", maxPages,
	)

	if bench {
		scraper.BenchmarkScraper(baseURL, dirPath, repeats, maxPages)
	} else {
		fmt.Println("- maxWorkers", maxWorkers)
		visited = scraper.New().Start(baseURL, dirPath, maxWorkers, maxPages)
		fmt.Println("\n[RESULT]",
			"\n- visited:", visited,
		)
	}
}
