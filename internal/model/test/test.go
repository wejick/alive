package test

// Test representing testcase in alive
type Test struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	Desc               string `json:"desc"`
	Domain             string `json:"domain"`
	Endpoint           string `json:"endpoint"`
	Method             string `json:"method"`
	Protocol           string `json:"protocol"`
	PeriodInCron       string `json:"period_in_cron"`
	Body               string `json:"body"`
	Header             string `json:"headers"`
	Agent              string `json:"agent"`
	ExpectedStatusCode int    `json:"expected_status_code"`
	Status             int    `json:"status"`
}
