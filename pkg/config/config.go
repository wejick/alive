package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Tests []Test `json:"tests"`
}
type Test struct {
	Name               string            `json:"name"`
	Desc               string            `json:"desc"`
	Domain             string            `json:"domain"`
	Endpoint           string            `json:"endpoint"`
	Method             string            `json:"method"`
	Protocol           string            `json:"protocol"`
	PeriodInCron       string            `json:"period_in_cron"`
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
