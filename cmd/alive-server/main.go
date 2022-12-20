package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/robfig/cron/v3"

	endpointAgent "github.com/wejick/alive/internal/endpoint/agent"
	repoAgent "github.com/wejick/alive/internal/repo/agent"
	serviceAgent "github.com/wejick/alive/internal/service/agent"

	httpTest "github.com/wejick/alive/internal/endpoint/test"
	repoTest "github.com/wejick/alive/internal/repo/test"
	serviceTest "github.com/wejick/alive/internal/service/test"

	"github.com/wejick/alive/pkg/metric"

	_ "modernc.org/sqlite"
)

var flagSqliteDBPath = flag.String("dbpath", "./alive.db", "path to sqlite db")

func main() {
	flag.Parse()

	// open sqlitedb
	fmt.Println("Opening config file", *flagSqliteDBPath)
	sqldb, err := sql.Open("sqlite", *flagSqliteDBPath+"?_pragma=foreign_keys(1)&_pragma=busy_timeout(1000)")
	if err != nil {
		fmt.Print(err)
		return
	}
	err = sqldb.Ping()
	if err != nil {
		fmt.Print(err)
		return
	}
	defer sqldb.Close()

	agentRepo := repoAgent.NewSqlite(sqldb)
	agentService := serviceAgent.New(agentRepo)
	agentHttpHandler := endpointAgent.New(agentService)
	healthcheckHandler := endpointAgent.NewHealthcheckWorker(agentService)

	testRepo := repoTest.NewSqlite(sqldb)
	testService := serviceTest.New(testRepo)
	testHttpHandler := httpTest.New(testService)

	metricRuntime := metric.New()
	metricRuntime.Init()

	c := cron.New()
	c.AddFunc("@every 20m", func() {
		//TODO log the error
		healthcheckHandler.UpdateHealthStatus()
	})
	c.Start()

	router := httprouter.New()
	router.GET("/config/agent", agentHttpHandler.GetAgentHandler)
	router.PUT("/config/agent", agentHttpHandler.AddAgentHandler)
	router.GET("/config/test", testHttpHandler.GetTestHandler)
	router.PUT("/config/test", testHttpHandler.AddTestHandler)

	router.GET("/agent/ping", agentHttpHandler.PingAgentHandler)

	fmt.Println("Alive serving 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
