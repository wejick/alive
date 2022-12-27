package agent

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wejick/alive/pkg/httputil"

	"github.com/julienschmidt/httprouter"
	model "github.com/wejick/alive/internal/model"
	serviceAgent "github.com/wejick/alive/internal/service/agent"
)

type Agent struct {
	AgentService *serviceAgent.Agent
}

type AgentHttpResponse struct {
	AgentList []model.Agent `json:"agent_list"`
}

func New(service *serviceAgent.Agent) *Agent {
	return &Agent{
		AgentService: service,
	}
}

// GetAgentHandler http handler for get agent
// param: id = list of id to get. empty means get all agents
func (A *Agent) GetAgentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idstring := r.URL.Query().Get("id")
	ids := strings.Split(idstring, ",")

	agentResponse := AgentHttpResponse{}
	agentResponse.AgentList = A.AgentService.GetAgents(ids)

	_ = httputil.ResponseJSON(agentResponse, 200, w)
}

// AddAgentHandler http handler for add agent
// param: json of agent
func (A *Agent) AddAgentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	agentParam := model.Agent{}
	err := json.NewDecoder(r.Body).Decode(&agentParam)
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusBadRequest, w)
		return
	}

	err = A.AgentService.AddAgent(agentParam)

	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
	} else {
		_ = httputil.ResponseError("", http.StatusAccepted, w)
	}
}

// PingAgentHandler http handler for agent to ping
func (A *Agent) PingAgentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := r.URL.Query().Get("id")
	err := A.AgentService.Ping(id)
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
	} else {
		_ = httputil.ResponseError("", http.StatusAccepted, w)
	}
}
