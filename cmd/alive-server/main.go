package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	httpAgent "github.com/wejick/alive/internal/endpoint/agent"
	repoAgent "github.com/wejick/alive/internal/repo/agent"
	serviceAgent "github.com/wejick/alive/internal/service/agent"
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

	metricRuntime := metric.New()
	metricRuntime.Init()

	router := httprouter.New()
	router.GET("/config/agent", agentHttpHandler.GetAgentHandler)
	router.PUT("/config/agent", agentHttpHandler.AddAgentHandler)
	router.POST("/agentping", nil)

	log.Fatal(http.ListenAndServe(":8081", router))
}
