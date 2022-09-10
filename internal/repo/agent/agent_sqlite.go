package agent

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	model "github.com/wejick/alive/internal/model"
	_ "modernc.org/sqlite"
)

// AgentSqlite ...
type AgentSqlite struct {
	db *sql.DB
}

func NewSqlite(sqliteDB *sql.DB) (agent *AgentSqlite) {
	agent = &AgentSqlite{db: sqliteDB}

	return
}

// GetAgents get 1 or more agent by id, empty array means get all agent
func (A *AgentSqlite) GetAgents(agentIDs []string) (agentList []model.Agent) {
	agentList = []model.Agent{}
	query := ""

	idcommaseparated := strings.Join(agentIDs[:], ",")

	if len(agentIDs) > 0 && agentIDs[0] != "" {
		query = fmt.Sprintf("SELECT * FROM agent where id in (%s)", idcommaseparated)
	} else {
		query = "SELECT * FROM agent"
	}

	// TODO : use prepared statement to avoid sql injection
	rows, err := A.db.Query(query)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		agent := model.Agent{}
		rows.Scan(&agent.ID, &agent.Location, &agent.GeoHash, &agent.ISP)
		agentList = append(agentList, agent)
	}

	return
}

// AddAgent add 1 agent at a time to db
func (A *AgentSqlite) AddAgent(agent model.Agent) (err error) {
	query := "INSERT INTO agent(location, geohash, ISP) VALUES(?,?,?)"

	tx, err := A.db.BeginTx(context.Background(), nil)
	if err != nil {
		return
	}
	_, err = tx.Exec(query, agent.Location, agent.GeoHash, agent.ISP)
	tx.Commit()

	return
}
