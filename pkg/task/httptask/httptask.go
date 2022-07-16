package httptask

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/wejick/alive/pkg/config"
	"github.com/wejick/alive/pkg/metric"
	"github.com/wejick/alive/pkg/task"
)

type HttpTask struct {
	config.Test
	task.Property
	httpclient *httpclient.Client
	header     http.Header
}

func NewHttpTask(agentcfg config.Agent, test config.Test, client *httpclient.Client, metricRuntime *metric.Metric) *HttpTask {
	return &HttpTask{
		Test:       test,
		httpclient: client,
		header:     constructHeader(test.Header),
		Property: task.Property{
			Location:      agentcfg.Location,
			ISP:           agentcfg.ISP,
			MetricRuntime: metricRuntime,
		},
	}
}

func (T *HttpTask) Run() {
	// setup
	var req *http.Request
	var res *http.Response

	// metric setup
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		T.MetricRuntime.MeasureAPILatency(duration, T.Name, T.Domain, T.Method, T.Result, res.Status, T.Location, T.ISP)
	}()

	// run the request and test
	var errNewRequest error
	switch T.Method {
	case "DELETE":
		fallthrough
	case "GET":
		req, errNewRequest = http.NewRequest(T.Method, constructURL(T.Protocol, T.Domain, T.Endpoint), nil)
		req.Header = T.header
	case "PATCH":
		fallthrough
	case "PUT":
		fallthrough
	case "POST":
		bodyBuffer := bytes.NewBufferString(T.Body)
		req, errNewRequest = http.NewRequest(T.Method, constructURL(T.Protocol, T.Domain, T.Endpoint), bodyBuffer)
		req.Header = T.header
	}

	if errNewRequest != nil {
		fmt.Println("couldn't create request ", T.Name)
		return
	}

	res, errDo := T.httpclient.Do(req)
	if errDo != nil {
		fmt.Println("Failed to request : ", errDo)
		return
	}

	fmt.Println(">> Run HttpTask : ", T.Test.Name, "WorkerID : ", T.Property.WorkerID)

	if res.StatusCode != T.Test.ExpectedStatusCode {
		T.Result = "FAILED"
		fmt.Println("Failed, expected status code : ", T.Test.ExpectedStatusCode, " got : ", res.StatusCode)
		return
	}

	T.Result = "PASS"
	fmt.Println("status code : ", res.Status)
}

const protocolshebang = "://"

func constructURL(protocol, domain, endpoint string) (url string) {
	return protocol + protocolshebang + domain + endpoint
}

func constructHeader(headersString map[string]string) (header http.Header) {
	if headersString == nil {
		return
	}
	header = http.Header{}
	for k, v := range headersString {
		header.Add(k, v)
	}
	return
}
