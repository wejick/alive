package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// StandardError response
type StandardError struct {
	Header *Header `json:"header,omitempty"`
}

type Page struct {
	TotalData   int64 `json:"total_data"`
	TotalPage   int64 `json:"total_page"`
	TotalInPage int64 `json:"total_in_page"`
	Page        int64 `json:"page"`
}

type Header struct {
	Message string `json:"status"`
}

type StandardResponse struct {
	Header *Header     `json:"header,omitempty"`
	Page   *Page       `json:"page,omitempty"`
	Data   interface{} `json:"data"`
}

// ResponseJSON response http request with application/json
func ResponseJSON(data interface{}, status int, writer http.ResponseWriter) (err error) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(status)

	response := StandardResponse{
		Data: data,
	}

	d, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		d, _ = json.Marshal(StandardError{Header: &Header{Message: "ResponseJSON: Failed to response" + err.Error()}})
		err = fmt.Errorf("ResponseJSON: Failed to response : %s", err)
	}

	_, errW := writer.Write(d)
	if errW != nil {
		fmt.Println("ResponseJSON: Couldn't write http response" + errW.Error())
	}
	return
}

// ResponseJsonPage response http request with application/json with page metadata
func ResponseJsonPage(data interface{}, message string, status int, page Page, writer http.ResponseWriter) (err error) {
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(status)

	response := StandardResponse{
		Header: &Header{
			Message: message,
		},
		Data: data,
		Page: &page,
	}

	d, err := json.Marshal(response)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		d, _ = json.Marshal(StandardError{Header: &Header{Message: "ResponseJSON: Failed to response" + err.Error()}})
		err = fmt.Errorf("ResponseJSON: Failed to response : %s", err)
	}

	_, errW := writer.Write(d)
	if errW != nil {
		fmt.Println("ResponseJSON: Couldn't write http response" + errW.Error())
	}
	return
}

// ResponseError response http request with standard error
func ResponseError(message string, status int, writer http.ResponseWriter) (err error) {
	return ResponseJSON(StandardError{Header: &Header{Message: message}}, status, writer)
}
