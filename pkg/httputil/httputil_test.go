package httputil

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseJSON(t *testing.T) {
	// null status ok
	handler := func(w http.ResponseWriter, r *http.Request) {
		_ = ResponseJSON(nil, http.StatusOK, w)
	}
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	assert.Equal(t, body, []byte("{\"data\":null}"))
	assert.Equal(t, w.Header(), http.Header(http.Header{"Content-Type": []string{"application/json"}}))
	assert.Equal(t, w.Code, 200)

	// interface status 500
	handler = func(w http.ResponseWriter, r *http.Request) {
		_ = ResponseError("Not OK", http.StatusInternalServerError, w)
	}
	w = httptest.NewRecorder()
	handler(w, req)
	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	assert.Equal(t, body, []byte(`{"data":{"header":{"status":"Not OK"}}}`))
	assert.Equal(t, w.Header(), http.Header(http.Header{"Content-Type": []string{"application/json"}}))
	assert.Equal(t, w.Code, 500)
}
