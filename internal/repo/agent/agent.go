package agent

import model "github.com/wejick/alive/internal/model"

type IAgent interface {
	GetAgents(agentIDs []string) []model.Agent
	GetAgentIDToSetUnhealthy(lastPingThreshold int) ([]int64, error)
	GetAgentIDToSetActive(lastPingThreshold int) ([]int64, error)

	AddAgent(agent model.Agent) error
	SetAgentStatus(agentIDs []string, status model.AgentStatus) error
	Ping(agentID string) error
}
