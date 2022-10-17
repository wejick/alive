package agent

import (
	"fmt"
	"strconv"
	"time"

	"github.com/wejick/alive/internal/model"
)

func (A *Agent) UpdateHealthStatus() (err error) {
	halfHourAgo := time.Now().Add(-time.Minute * 30)

	// get agent to set active
	toSetActive, err := A.agentRepo.GetAgentIDToSetActive(int(halfHourAgo.Unix()))
	if err != nil {
		return fmt.Errorf("UpdateHealthStatus: %s", err.Error())
	}

	// get agent to set unhealthy
	toSetUnhealthy, err := A.agentRepo.GetAgentIDToSetUnhealthy(int(halfHourAgo.Unix()))
	if err != nil {
		return fmt.Errorf("UpdateHealthStatus: %s", err.Error())
	}

	// set agent status to active
	err = A.agentRepo.SetAgentStatus(arrayInt2arrayString(toSetActive), model.AgentStatusActive)
	if err != nil {
		return fmt.Errorf("UpdateHealthStatus: %s", err.Error())
	}

	// set agent status to unhealty
	err = A.agentRepo.SetAgentStatus(arrayInt2arrayString(toSetUnhealthy), model.AgentStatusUnhealty)
	if err != nil {
		return fmt.Errorf("UpdateHealthStatus: %s", err.Error())
	}

	return
}

func arrayInt2arrayString(input []int64) (output []string) {
	output = []string{}
	for _, item := range input {
		output = append(output, strconv.Itoa(int(item)))
	}
	return
}

// belum di test
