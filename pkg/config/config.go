package config

import (
	"encoding/json"
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
	agentName     string
}

type Config struct {
	ServerAddress string `json:"server-address"`
	Agent         Agent  `json:"agent"`
	Tests         []Test `json:"tests"`
}

type Agent struct {
	Location string `json:"location"`
	GeoHash  string `json:"geohash"`
	ISP      string `json:"ISP"`
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

func NewHTTPConfigLoader(serverAddress string, agentName string) (loader *ConfigLoader) {
	// Create a new HTTP client with a default timeout
	timeout := 15000 * time.Millisecond
	httpclt := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	loader = &ConfigLoader{
		httpclient:    httpclt,
		serverAddress: serverAddress,
		agentName:     agentName,
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
	config.Agent, err = C.GetAgentConfig(C.agentName)
	if err != nil {
		return
	}
	config.Tests, err = C.GetTestConfigByAgent(C.agentName)
	if err != nil {
		return
	}
	return
}

// GetTestConfigByClient getting test config from server
func (C *ConfigLoader) GetTestConfigByAgent(AgentID string) (testConfig []Test, err error) {
	testResp := testEndpoint.TestHttpResponse{}
	resp, err := C.httpclient.Get(C.serverAddress+testConfigPath+"?agentid="+AgentID, nil)
	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&testResp)
	testConfig = modeltoTestConfig(testResp.TestList)
	return
}

// GetTestConfigByClient getting test config from server
func (C *ConfigLoader) GetAgentConfig(AgentID string) (agentConfig Agent, err error) {
	agentResponse := agentEndpoint.AgentHttpResponse{}
	resp, err := C.httpclient.Get(C.serverAddress+agentConfigPath+"?id="+AgentID, nil)
	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&agentResponse)
	agentConfig = modeltoAgentConfig(agentResponse.AgentList)
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
	agentConfig.Location = agentList[0].Location
	agentConfig.GeoHash = agentList[0].GeoHash
	agentConfig.ISP = agentList[0].ISP

	return
}
