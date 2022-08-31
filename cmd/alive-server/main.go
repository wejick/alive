package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	httpAgent "github.com/wejick/alive/internal/endpoint/agent"
	repoAgent "github.com/wejick/alive/internal/repo/agent"
	serviceAgent "github.com/wejick/alive/internal/service/agent"

	httpTest "github.com/wejick/alive/internal/endpoint/test"
	repoTest "github.com/wejick/alive/internal/repo/test"
	serviceTest "github.com/wejick/alive/internal/service/test"

	"github.com/wejick/alive/pkg/metric"

	_ "modernc.org/sqlite"
)

func main() {
	// open sqlitedb
	sqldb, err := sql.Open("sqlite", "./alive.db")
	if err != nil {
		return
	}
	defer sqldb.Close()

	agentRepo := repoAgent.NewSqlite(sqldb)
	agentService := serviceAgent.New(agentRepo)
	agentHttpHandler := httpAgent.New(agentService)

	testRepo := repoTest.NewSqlite(sqldb)
	testService := serviceTest.New(testRepo)
	testHttpHandler := httpTest.New(testService)

	metricRuntime := metric.New()
	metricRuntime.Init()

	router := httprouter.New()
	router.GET("/config/agent", agentHttpHandler.GetAgentHandler)
	router.PUT("/config/agent", agentHttpHandler.AddAgentHandler)
	router.GET("/config/test", testHttpHandler.GetTestHandler)
	router.PUT("/config/test", testHttpHandler.AddTestHandler)

	router.POST("/agentping", nil)

	fmt.Println("Alive serving 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
