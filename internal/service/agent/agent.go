package agent

import (
	model "github.com/wejick/alive/internal/model"
	repoAgent "github.com/wejick/alive/internal/repo/agent"
)

type IAgent interface {
	AddAgent(model.Agent) error
	GetAgents([]string) []model.Agent
	Ping(string) error
	UpdateHealthStatus() error
}

type Agent struct {
	agentRepo repoAgent.IAgent
}

func New(agentRepo repoAgent.IAgent) *Agent {
	return &Agent{
		agentRepo: agentRepo,
	}
}

func (A *Agent) GetAgents(agentIDs []string) (agentList []model.Agent) {
	return A.agentRepo.GetAgents(agentIDs)
}

func (A *Agent) AddAgent(agent model.Agent) (err error) {
	return A.agentRepo.AddAgent(agent)
}

func (A *Agent) Ping(agentID string) (err error) {
	return A.agentRepo.Ping(agentID)
}
