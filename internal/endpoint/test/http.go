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
	agentStr := r.URL.Query().Get("agentid")

	rows, _ := strconv.Atoi(rowsStr)
	page, _ := strconv.Atoi(pageStr)

	resp := TestHttpResponse{}
	var err error
	resp.TestList, err = T.testService.GetTest(ids, agentStr, int64(rows), int64(page))
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	totalTest, err := T.testService.GetTotalTest()
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
		return
	}

	totalPage := int64(1)
	if rows > 0 {
		totalPage = totalTest / int64(rows)
	}
	pageInfo := httputil.Page{
		TotalData:   totalTest,
		TotalPage:   totalPage,
		TotalInPage: int64(len(resp.TestList)),
		Page:        int64(page),
	}
	_ = httputil.ResponseJsonPage(resp, "", http.StatusOK, pageInfo, w)
}

func (T *Test) AddTestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	testParam := model.Test{}
	err := json.NewDecoder(r.Body).Decode(&testParam)
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusBadRequest, w)
		return
	}

	err = T.testService.AddTest(testParam)
	if err != nil {
		_ = httputil.ResponseError(err.Error(), http.StatusInternalServerError, w)
	} else {
		_ = httputil.ResponseError("", http.StatusAccepted, w)
	}
}
