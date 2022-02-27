package httptask

import (
	"fmt"
	"net/http"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/wejick/alive/pkg/config"
	"github.com/wejick/alive/pkg/task"
)

type HttpTask struct {
	config.Test
	task.Property
	httpclient *httpclient.Client
	header     http.Header
}

func NewHttpTask(test config.Test, client *httpclient.Client) *HttpTask {
	return &HttpTask{
		Test:       test,
		httpclient: client,
		header:     constructHeader(test.Header),
	}
}

func (T *HttpTask) Run() {
	var req *http.Request
	var errNewRequest error
	switch T.Method {
	case "GET":
		req, errNewRequest = http.NewRequest(T.Method, constructURL(T.Protocol, T.Domain, T.Endpoint), nil)
		req.Header = T.header
	case "POST":

	}

	if errNewRequest != nil {
		fmt.Println("couldn't create request ", T.Name)
		return
	}

	res, _ := T.httpclient.Do(req)

	fmt.Println(">> Run HttpTask : ", T.Test.Name, "WorkerID : ", T.Property.WorkerID)
	fmt.Println("Headers : ", T.header)

	if res.StatusCode != T.Test.ExpectedStatusCode {
		fmt.Println("Failed, expected status code : ", T.Test.ExpectedStatusCode, " got : ", res.StatusCode)
		return
	}
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
