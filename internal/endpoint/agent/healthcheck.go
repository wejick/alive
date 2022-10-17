package agent

import serviceAgent "github.com/wejick/alive/internal/service/agent"

// HealthcheckWorker ...
type HealthcheckWorker struct {
	AgentService *serviceAgent.Agent
}

func NewHealthcheckWorker(service *serviceAgent.Agent) *HealthcheckWorker {
	return &HealthcheckWorker{
		AgentService: service,
	}
}

func (H *HealthcheckWorker) UpdateHealthStatus() (err error) {
	err = H.AgentService.UpdateHealthStatus()
	return
}
