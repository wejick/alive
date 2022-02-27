package task

import "github.com/robfig/cron/v3"

type Property struct {
	ID       string
	Status   int
	WorkerID cron.EntryID
}

const (
	StatusRunning = iota
	StatusPaused
)
