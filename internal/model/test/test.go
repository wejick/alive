package test

// Test is representing test data of API monitoring
type Test struct {
	ID                 string            `json:"id"`
	Name               string            `json:"name"`
	Desc               string            `json:"desc"`
	Domain             string            `json:"domain"`
	Endpoint           string            `json:"endpoint"`
	Method             string            `json:"method"`
	Protocol           string            `json:"protocol"`
	PeriodInCron       string            `json:"period_in_cron"`
	Body               string            `json:"body"`
	Header             map[string]string `json:"headers"`
	Agent              []string          `json:"agents"`
	ExpectedStatusCode int               `json:"expected_status_code"`
	Status             bool              `json:"status"`
}
