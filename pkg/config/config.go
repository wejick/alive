package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	testEndpoint "github.com/wejick/alive/internal/endpoint/test"
	model "github.com/wejick/alive/internal/model"

	agentEndpoint "github.com/wejick/alive/internal/endpoint/agent"
)

type ConfigLoader struct {
	httpclient    *httpclient.Client
	serverAddress string
	agentID       int64
}

type Config struct {
	BaseID        int64  `json:"agent-id"`
	ServerAddress string `json:"server-address"`
	Agent         Agent  `json:"agent"`
	Tests         []Test `json:"tests"`
}

type Agent struct {
	ID         int64  `json:"id"`
	Location   string `json:"location"`
	GeoHash    string `json:"geohash"`
	ISP        string `json:"ISP"`
	StatusText string `json:"status_text"`
	Status     int    `json:"status"`
}

type Test struct {
	Name               string            `json:"name"`
	Desc               string            `json:"desc"`
	Domain             string            `json:"domain"`
	Endpoint           string            `json:"endpoint"`
	Method             string            `json:"method"`
	Protocol           string            `json:"protocol"`
	PeriodInCron       string            `json:"period_in_cron"`
	Body               string            `json:"body"`
	Header             map[string]string `json:"headers"`
	ExpectedStatusCode int               `json:"expected_status_code"`
}

const agentConfigPath = "/config/agent"
const testConfigPath = "/config/test"
const pingPath = "/agent/ping"

func NewHTTPConfigLoader(serverAddress string, agentID int64) (loader *ConfigLoader) {
	// Create a new HTTP client with a default timeout
	timeout := 15000 * time.Millisecond
	httpclt := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	loader = &ConfigLoader{
		httpclient:    httpclt,
		serverAddress: serverAddress,
		agentID:       agentID,
	}

	return
}

// LoadConfig load config
func LoadConfig(path string) (config Config, err error) {
	F, err := os.Open(path)
	if err != nil {
		return
	}
	defer F.Close()

	jsonDecoder := json.NewDecoder(F)
	jsonDecoder.Decode(&config)

	return
}

// GetConfigFromServer getting whole config from server
func (C *ConfigLoader) GetConfigFromServer() (config Config, err error) {
	config.Agent, err = C.GetAgentConfig(C.agentID)
	if err != nil {
		return
	}
	config.Tests, err = C.GetTestConfigByAgent(C.agentID)
	if err != nil {
		return
	}
	return
}

// GetTestConfigByClient getting test config from server
func (C *ConfigLoader) GetTestConfigByAgent(AgentID int64) (testConfig []Test, err error) {
	path := fmt.Sprintf("%s?id=%d", C.serverAddress+testConfigPath, C.agentID)

	testResp := struct {
		Data testEndpoint.TestHttpResponse `json:"data"`
	}{}
	resp, err := C.httpclient.Get(path, nil)
	if err != nil {
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&testResp)
	if err != nil {
		return
	}
	testConfig = modeltoTestConfig(testResp.Data.TestList)
	return
}

// GetTestConfigByClient getting test config from server
func (C *ConfigLoader) GetAgentConfig(AgentID int64) (agentConfig Agent, err error) {
	path := fmt.Sprintf("%s?id=%d", C.serverAddress+agentConfigPath, C.agentID)

	agentResponse := struct {
		Data agentEndpoint.AgentHttpResponse `json:"data"`
	}{}
	resp, err := C.httpclient.Get(path, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&agentResponse)
	if err != nil {
		return
	}

	agentConfig = modeltoAgentConfig(agentResponse.Data.AgentList)

	return
}

func (C *ConfigLoader) Ping() (err error) {
	path := fmt.Sprintf("%s?id=%d", C.serverAddress+pingPath, C.agentID)
	_, err = C.httpclient.Get(path, nil)
	return
}

func modeltoTestConfig(testList []model.Test) (testConfig []Test) {
	for idx, testM := range testList {
		testConfig = append(testConfig, Test{
			Name:               testM.Name,
			Desc:               testM.Desc,
			Domain:             testM.Domain,
			Endpoint:           testM.Endpoint,
			Method:             testM.Method,
			Protocol:           testM.Protocol,
			PeriodInCron:       testM.PeriodInCron,
			Body:               testM.Body,
			ExpectedStatusCode: testM.ExpectedStatusCode,
		})
		json.Unmarshal([]byte(testM.Header), &testConfig[idx].Header)
	}
	return
}

func modeltoAgentConfig(agentList []model.Agent) (agentConfig Agent) {
	if len(agentList) == 0 {
		return
	}
	agentConfig.ID = agentList[0].ID
	agentConfig.Location = agentList[0].Location
	agentConfig.GeoHash = agentList[0].GeoHash
	agentConfig.ISP = agentList[0].ISP
	agentConfig.StatusText = agentList[0].StatusText
	agentConfig.Status = agentList[0].Status

	return
}
