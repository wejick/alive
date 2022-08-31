package agent

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/wejick/alive/pkg/httputil"

	"github.com/julienschmidt/httprouter"
	modelAgent "github.com/wejick/alive/internal/model/agent"
	serviceAgent "github.com/wejick/alive/internal/service/agent"
)

type Agent struct {
	AgentService *serviceAgent.Agent
}

type AgentHttpResponse struct {
	AgentList []modelAgent.Agent `json:"agent_list"`
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

	httputil.ResponseJSON(agentResponse, 200, w)
}

// AddAgentHandler http handler for add agent
// param: json of agent
func (A *Agent) AddAgentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	agentParam := modelAgent.Agent{}
	json.NewDecoder(r.Body).Decode(&agentParam)

	err := A.AgentService.AddAgent(agentParam)

	if err != nil {
		httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
	} else {
		httputil.ResponseError("", http.StatusAccepted, w)
	}
}
