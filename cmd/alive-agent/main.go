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
	fmt.Println("Local config loaded")

	cfgLoader := config.NewHTTPConfigLoader(globalConfig.ServerAddress, globalConfig.BaseID)

	httpConfigLoaderFunc := func() {
		conf, err := cfgLoader.GetConfigFromServer()
		if err != nil {
			return
		}
		globalConfig.Agent = conf.Agent
		globalConfig.Tests = conf.Tests

		fmt.Println("Remote config fetched from :", globalConfig.ServerAddress)
	}
	httpPingFunc := func() {
		cfgLoader.Ping()
	}

	httpConfigLoaderFunc()
	httpPingFunc()
	c := cron.New()
	c.AddFunc("@every 30s", httpConfigLoaderFunc)
	c.AddFunc("@every 30s", httpPingFunc)
	c.Start()

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

	StartUpMessage(globalConfig)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func constructTaskList(config config.Config, client *httpclient.Client, metricRuntime *metric.Metric) (taskList []*httptask.HttpTask) {
	for _, test := range config.Tests {
		taskList = append(taskList, httptask.NewHttpTask(config.Agent, test, client, metricRuntime))
	}
	return
}

func StartUpMessage(config config.Config) {
	fmt.Printf(`Alive Agent is starting up
	Agent ID : %d
	Agent ISP : %s
	Agent Location : %s
	Agent GeoHash : %s
	
	Number of Test : %d`, config.BaseID, config.Agent.ISP, config.Agent.Location, config.Agent.GeoHash, len(config.Tests))
}
