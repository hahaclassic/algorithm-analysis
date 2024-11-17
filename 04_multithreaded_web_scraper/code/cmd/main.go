package main

import (
	"flag"

	"github.com/hahaclassic/algorithm-analysis/04_multithreaded_web_scrapper/code/internal/scrapper"
	"github.com/hahaclassic/algorithm-analysis/04_multithreaded_web_scrapper/code/pkg/logger"
)

var (
	baseURL    string
	dirPath    string
	maxWorkers int
	maxPages   int
)

func init() {
	flag.StringVar(&baseURL, "url", "https://edimdoma.ru", "the main page of the web resource")
	flag.StringVar(&dirPath, "dir", "../data", "the directory where the received html pages are saved")
	flag.IntVar(&maxWorkers, "workers", 1, "max number of goroutines")
	flag.IntVar(&maxPages, "pages", 10, "max number of pages")
	flag.Parse()
}

func init() {
	logger.SetupLogger("../log.txt")
}

func main() {
	scrapper.New().Start(baseURL, dirPath, maxWorkers, maxPages)
}
