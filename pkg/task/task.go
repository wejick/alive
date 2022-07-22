package task

import (
	"github.com/robfig/cron/v3"
	"github.com/wejick/alive/pkg/metric"
)

type Property struct {
	ID            string
	Location      string
	GeoHash       string
	ISP           string
	Result        string
	Status        int
	WorkerID      cron.EntryID
	MetricRuntime *metric.Metric
}

const (
	StatusRunning = iota
	StatusPaused
)
