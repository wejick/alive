package httptask

import (
	"fmt"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/wejick/alive/pkg/config"
)

type HttpTask struct {
	config.Test
	httpclient *httpclient.Client
}

func NewHttpTask(test config.Test, client *httpclient.Client) *HttpTask {
	return &HttpTask{
		Test:       test,
		httpclient: client,
	}
}

func (T *HttpTask) Run() {
	res, _ := T.httpclient.Get(T.Protocol+"://"+T.Domain+T.Endpoint, nil)

	fmt.Println("Run HttpTask : ", T.Test.Name)

	if res.StatusCode != T.Test.ExpectedStatusCode {
		fmt.Println("Failed, expected status code : ", T.Test.ExpectedStatusCode, " got : ", res.StatusCode)
		return
	}
	fmt.Println("status code : ", res.Status)
}
