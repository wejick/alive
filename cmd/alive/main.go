package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/robfig/cron/v3"
	"github.com/wejick/alive/pkg/config"
	"github.com/wejick/alive/pkg/task/httptask"
)

func main() {
	globalConfig, err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Println("couldn't open config", err)
		os.Exit(-1)
	}
	fmt.Println("Config loaded")

	// Create a new HTTP client with a default timeout
	timeout := 15000 * time.Millisecond
	httpclt := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	httpTaskList := constructTaskList(globalConfig.Tests, httpclt)

	worker := cron.New()
	for _, task := range httpTaskList {
		task.Property.WorkerID, _ = worker.AddJob(task.PeriodInCron, task)
	}
	worker.Start()

	select {}
}

func constructTaskList(tests []config.Test, client *httpclient.Client) (taskList []*httptask.HttpTask) {
	for _, test := range tests {
		taskList = append(taskList, httptask.NewHttpTask(test, client))
	}
	return
}
