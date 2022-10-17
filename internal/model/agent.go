package model

import "time"

// Agent is representing agent in the alive system
type Agent struct {
	ID         int64  `json:"id"`
	Location   string `json:"location"`
	GeoHash    string `json:"geohash"`
	ISP        string `json:"ISP"`
	StatusText string `json:"status_text"`
	Status     int    `json:"status"`
}

type AgentStatus int

const (
	AgentStatusInactive AgentStatus = 0
	AgentStatusActive   AgentStatus = 1
	AgentStatusStarting AgentStatus = 2
	AgentStatusUnhealty AgentStatus = 3
)

const (
	AgentStatusInactiveText = "INACTIVE"
	AgentStatusActiveText   = "ACTIVE"
	AgentStatusStartingText = "STARTING"
	AgentStatusUnhealtyText = "SUSPENDED"
)

func GetAgentStatusText(status AgentStatus) string {
	switch status {
	case AgentStatusInactive:
		return AgentStatusInactiveText
	case AgentStatusActive:
		return AgentStatusActiveText
	case AgentStatusStarting:
		return AgentStatusStartingText
	case AgentStatusUnhealty:
		return AgentStatusUnhealtyText
	}

	return ""
}

type AgentPing struct {
	LastPingTime       time.Time
	LastPingTimeString string `json:"last_ping_time"`
	AgentID            int    `json:"agent_id"`
}
