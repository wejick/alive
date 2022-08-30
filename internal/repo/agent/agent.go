package agent

import modelAgent "github.com/wejick/alive/internal/model/agent"

type IAgent interface {
	GetAgents(agentIDs ...string) []modelAgent.Agent
	AddAgent(agent modelAgent.Agent) error
}
