package test

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	model "github.com/wejick/alive/internal/model"
	serviceTest "github.com/wejick/alive/internal/service/test"
	"github.com/wejick/alive/pkg/httputil"
)

type Test struct {
	testService *serviceTest.Test
}

type TestHttpResponse struct {
	TestList []model.Test `json:"test_list"`
}

func New(testService *serviceTest.Test) *Test {
	return &Test{
		testService: testService,
	}
}

func (T *Test) GetTestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idstring := r.URL.Query().Get("id")
	ids := strings.Split(idstring, ",")

	rowsStr := r.URL.Query().Get("rows")
	pageStr := r.URL.Query().Get("page")
	agentStr := r.URL.Query().Get("agent")

	rows, _ := strconv.Atoi(rowsStr)
	page, _ := strconv.Atoi(pageStr)

	resp := TestHttpResponse{}
	var err error
	resp.TestList, err = T.testService.GetTest(ids, agentStr, int64(rows), int64(page))
	if err != nil {
		httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	httputil.ResponseJSON(resp, 200, w)
}

func (T *Test) AddTestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	testParam := model.Test{}
	json.NewDecoder(r.Body).Decode(&testParam)

	err := T.testService.AddTest(testParam)
	if err != nil {
		httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
	} else {
		httputil.ResponseError("", http.StatusAccepted, w)
	}
}
