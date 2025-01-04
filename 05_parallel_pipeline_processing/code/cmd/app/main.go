package main

import (
	"flag"
	"log"

	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/config"
	"github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/internal/app"
)

var (
	sourceDir       string
	chanLen         int
	stageWorkers    int
	clearCollection bool
)

func init() {
	flag.StringVar(&sourceDir, "dir", "../data", "the directory where the received html pages are saved")
	flag.IntVar(&chanLen, "chan", 1, "len of ")
	flag.IntVar(&stageWorkers, "workers", 1, "number of goroutines per stage")
	flag.BoolVar(&clearCollection, "d", false, "flag for delete all recipes in db")
	flag.Parse()

	if chanLen < 0 || stageWorkers < 1 {
		log.Fatal("invalid number of stage workers or channel length")
	}
}

func main() {
	conf := config.MustLoad()
	conf.ClearCollection = clearCollection
	conf.Pipeline = config.PipelineConfig{
		SourceDir:    sourceDir,
		ChanLen:      chanLen,
		StageWorkers: stageWorkers,
	}

	app.Run(conf)
}
