package models

import (
	"fmt"
	"time"

	tableoutput "github.com/hahaclassic/algorithm-analysis/05_parallel_pipeline/pkg/table"
	"github.com/jedib0t/go-pretty/table"
)

type AvgStageInfo struct {
	Name           string
	WaitingTime    time.Duration // Started - Queued
	ProcessingTime time.Duration // Finished - Started
	countTask      int
}

type AvgStatInfo struct {
	TaskLifeTime  time.Duration // lifeTime = Destucted - Created
	countTask     int
	StageInfo     []*AvgStageInfo
	minCreated    time.Time
	maxDestructed time.Time
}

func (s *AvgStatInfo) AddTask(task *Task) {
	s.countTask++
	s.TaskLifeTime += task.Destructed.Sub(task.Created)

	if s.minCreated.IsZero() || task.Created.Compare(s.minCreated) < 0 {
		s.minCreated = task.Created
	}
	if s.maxDestructed.IsZero() || task.Destructed.Compare(s.maxDestructed) > 0 {
		s.maxDestructed = task.Destructed
	}

	for i, stage := range task.Stages {
		if i >= len(s.StageInfo) {
			s.StageInfo = append(s.StageInfo, &AvgStageInfo{
				Name: stage.Name,
			})
		}

		s.StageInfo[i].countTask++
		s.StageInfo[i].WaitingTime = stage.Started.Sub(stage.Queued)
		s.StageInfo[i].ProcessingTime = stage.Finished.Sub(stage.Started)
	}
}

func (s *AvgStatInfo) Avg() {
	s.TaskLifeTime /= time.Duration(s.countTask)
	for name := range s.StageInfo {
		t := time.Duration(s.StageInfo[name].countTask)
		s.StageInfo[name].ProcessingTime /= t
		s.StageInfo[name].WaitingTime /= t
	}
}

func (s *AvgStatInfo) Print(workers int, chanLen int) {
	s.Avg()

	fmt.Println("--------------------------\nStatistics\n--------------------------")
	fmt.Println("Workers:", workers, "\nChannel length:", chanLen, "\nNum of tasks:",
		s.countTask, "\nAvg task lifetime:", s.TaskLifeTime, "\nTotal time:", s.maxDestructed.Sub(s.minCreated))

	headers := []string{"Stage Name", "Avg waiting time", "Avg processing time"}
	rows := make([][]interface{}, 0, len(s.StageInfo))
	for _, stage := range s.StageInfo {
		rows = append(rows, []interface{}{stage.Name, stage.WaitingTime, stage.ProcessingTime})
	}

	tableoutput.PrintTable(table.StyleDefault, headers, rows)
}
