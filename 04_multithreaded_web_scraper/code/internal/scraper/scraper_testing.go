package scraper

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func BenchmarkScraper(baseURL, saveDir string, numOfRepeats, maxPages int) {
	logicalCores := runtime.NumCPU()
	threadCounts := []int{1, 2, 4, 8, 16, 32, 4 * logicalCores}

	for _, numThreads := range threadCounts {
		var sumElapsed time.Duration
		var processedPages int
		for i := 0; i < numOfRepeats; i++ {
			fmt.Println("test ", i+1)
			scraper := New()
			startTime := time.Now()
			processedPages += scraper.Start(baseURL, saveDir, numThreads, maxPages)
			sumElapsed += time.Since(startTime)
			os.RemoveAll(saveDir)
		}

		pagesPerSecond := float64(processedPages) / float64(sumElapsed.Seconds())
		fmt.Printf("Threads: %d, Pages/sec: %.2f\n", numThreads, pagesPerSecond)
	}
}
