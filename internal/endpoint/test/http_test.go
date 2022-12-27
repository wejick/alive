package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/wejick/alive/internal/model"

	serviceTest "github.com/wejick/alive/internal/service/test"
	mockTestRepo "github.com/wejick/alive/mocks/internal_mock/repo/test"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestTest_GetTestHandler(t *testing.T) {
	testRepo := mockTestRepo.NewItest(t)
	testService := serviceTest.New(testRepo)
	httpHandler := New(testService)

	recorder := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	testRepo.On("GetTotalTest").Return(int64(0), nil)
	testRepo.On("GetTest", []string{""}, "", 0, 0).Return([]model.Test{}, nil)
	httpHandler.GetTestHandler(recorder, r, nil)

	res := recorder.Result()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, string(data), `{"header":{"status":""},"page":{"total_data":0,"total_page":1,"total_in_page":0,"page":0},"data":{"test_list":[]}}`)
}
