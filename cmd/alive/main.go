package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/robfig/cron/v3"
	"github.com/wejick/alive/pkg/config"
	"github.com/wejick/alive/pkg/metric"
	"github.com/wejick/alive/pkg/task/httptask"
)

func main() {
	globalConfig, err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Println("couldn't open config", err)
		os.Exit(-1)
	}
	fmt.Println("Config loaded")

	metricRuntime := metric.New()

	// Create a new HTTP client with a default timeout
	timeout := 15000 * time.Millisecond
	httpclt := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	httpTaskList := constructTaskList(globalConfig, httpclt, metricRuntime)

	worker := cron.New()
	for _, task := range httpTaskList {
		task.Property.WorkerID, _ = worker.AddJob(task.PeriodInCron, task)
	}
	worker.Start()

	metricRuntime.Init()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func constructTaskList(config config.Config, client *httpclient.Client, metricRuntime *metric.Metric) (taskList []*httptask.HttpTask) {
	for _, test := range config.Tests {
		taskList = append(taskList, httptask.NewHttpTask(config.Agent, test, client, metricRuntime))
	}
	return
}
