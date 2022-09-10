package agent

import model "github.com/wejick/alive/internal/model"

type IAgent interface {
	GetAgents(agentIDs []string) []model.Agent
	AddAgent(agent model.Agent) error
}
