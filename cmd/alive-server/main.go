package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/wejick/alive/pkg/config"
	"github.com/wejick/alive/pkg/metric"
)

func main() {
	globalConfig, err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Println("couldn't open config", err)
		os.Exit(-1)
	}
	fmt.Println("Config loaded")

	metricRuntime := metric.New()
	metricRuntime.Init()

	router := httprouter.New()
	router.GET("/config/:agentID", nil)
	router.POST("/agentping", nil)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
