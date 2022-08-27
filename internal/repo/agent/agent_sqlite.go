package agent

import (
	"database/sql"
	"fmt"

	modelAgent "github.com/wejick/alive/internal/model/agent"

	_ "modernc.org/sqlite"
)

// AgentSqlite ...
type AgentSqlite struct {
	db *sql.DB
}

func New(filePath string, sqliteDB *sql.DB) (agent *AgentSqlite) {
	agent = &AgentSqlite{db: sqliteDB}

	return
}

// GetAgents get 1 or more agent by id, empty array means get all agent
func (A *AgentSqlite) GetAgents(agentIDs ...string) (agentList []modelAgent.Agent) {
	agentList = []modelAgent.Agent{}
	query := ""

	idcommaseparated := ""
	for idx, id := range agentIDs {
		if idx == 0 {
			idcommaseparated = id
		} else {
			idcommaseparated = id + "," + idcommaseparated
		}
	}

	if len(agentIDs) > 0 {
		query = fmt.Sprintf("SELECT * FROM agent where id in (%s)", idcommaseparated)
	} else {
		query = "SELECT * FROM agent"
	}

	// TODO : use prepared statement to avoid sql injection
	rows, err := A.db.Query(query)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	defer rows.Close()
	if rows.Next() {
		agent := modelAgent.Agent{}
		rows.Scan(&agent.ID, &agent.Location, &agent.GeoHash, &agent.ISP)
		agentList = append(agentList, agent)
	}

	return
}