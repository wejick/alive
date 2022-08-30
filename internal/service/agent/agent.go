package agent

import (
	modelAgent "github.com/wejick/alive/internal/model/agent"
	repoAgent "github.com/wejick/alive/internal/repo/agent"
)

type Agent struct {
	agentRepo repoAgent.IAgent
}

func New(agentRepo repoAgent.IAgent) *Agent {
	return &Agent{
		agentRepo: agentRepo,
	}
}

func (A *Agent) GetAgents(agentIDs ...string) (agentList []modelAgent.Agent) {
	return A.agentRepo.GetAgents(agentIDs...)
}

func (A *Agent) AddAgent(agent modelAgent.Agent) (err error) {
	return A.agentRepo.AddAgent(agent)
}
