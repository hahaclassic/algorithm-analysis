package models

import (
	"time"
)

type Task struct {
	ID         int
	Stages     []*StageInfo // Pipeline stages
	Created    time.Time    // The time when the task was created
	Destructed time.Time    // The time when the task was destructed
	Meta       any          // data
}

type StageInfo struct {
	Name     string
	Queued   time.Time
	Started  time.Time
	Finished time.Time
}
