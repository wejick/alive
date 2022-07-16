package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Agent Agent  `json:"agent"`
	Tests []Test `json:"tests"`
}

type Agent struct {
	Location string `json:"location"`
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
