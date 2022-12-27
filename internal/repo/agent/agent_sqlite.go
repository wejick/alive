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
		rows.Scan(&agent.ID, &agent.Location, &agent.GeoHash, &agent.ISP, &agent.Status)
		agent.StatusText = model.GetAgentStatusText(model.AgentStatus(agent.Status))
		agentList = append(agentList, agent)
	}

	return
}

// AddAgent add 1 agent at a time to db
func (A *AgentSqlite) AddAgent(agent model.Agent) (err error) {
	query := "INSERT INTO agent(location, geohash, ISP, status) VALUES(?,?,?,?)"

	tx, err := A.db.BeginTx(context.Background(), nil)
	if err != nil {
		return
	}
	_, err = tx.Exec(query, agent.Location, agent.GeoHash, agent.ISP, agent.Status)
	if err != nil {
		return
	}
	err = tx.Commit()

	return
}

func (A *AgentSqlite) SetAgentStatus(agentIDs []string, status model.AgentStatus) (err error) {
	query := "UPDATE agent SET status=? where id in (?)"
	idcommaseparated := strings.Join(agentIDs[:], ",")

	tx, err := A.db.BeginTx(context.Background(), nil)
	if err != nil {
		return
	}
	_, err = tx.Exec(query, status, idcommaseparated)
	if err != nil {
		return
	}
	err = tx.Commit()

	return
}

func (A *AgentSqlite) Ping(agentID string) (err error) {
	query := "INSERT INTO agent_ping(agent_id,last_ping_time) VALUES(?, strftime('%s','now')) ON CONFLICT(agent_id) DO UPDATE SET last_ping_time=strftime('%s','now')"
	tx, err := A.db.BeginTx(context.Background(), nil)
	if err != nil {
		return
	}
	_, err = tx.Exec(query, agentID)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}

// GetAgentIDToSetUnhealthy get agent with status active but have last ping more than lastPingThreshold
// lastPingThreshold in unix time
func (A *AgentSqlite) GetAgentIDToSetUnhealthy(lastPingThreshold int) (ids []int64, err error) {
	query := "SELECT * FROM agent WHERE id = (SELECT id FROM agent_ping WHERE agent_ping.last_ping_time < ?) AND agent.status = ?"
	rows, err := A.db.Query(query, lastPingThreshold, model.AgentStatusActive)
	if err != nil {
		return ids, fmt.Errorf("GetAgentIDByStatus : %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var agentID int64
		err = rows.Scan(&agentID)
		if err != nil {
			break
		}
		ids = append(ids, agentID)
	}

	return
}

// GetAgentIDToSetActive get agent with status unhealthy but have last ping less than lastPingThreshold
// lastPingThreshold in unix time
func (A *AgentSqlite) GetAgentIDToSetActive(lastPingThreshold int) (ids []int64, err error) {
	query := "SELECT * FROM agent WHERE id = (SELECT id FROM agent_ping WHERE agent_ping.last_ping_time > ?) AND agent.status = ?"
	rows, err := A.db.Query(query, lastPingThreshold, model.AgentStatusUnhealty)
	if err != nil {
		return ids, fmt.Errorf("GetAgentIDByStatus : %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var agentID int64
		err = rows.Scan(&agentID)
		if err != nil {
			break
		}
		ids = append(ids, agentID)
	}

	return
}
