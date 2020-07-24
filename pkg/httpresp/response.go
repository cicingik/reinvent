package httpresp

import (
	"encoding/json"
	"net/http"
)

const (
	ContentTypeJson = "application/json; charset=UTF-8"
)

type (
	HTTPResponseWrapper struct {
		HttpCode int         `json:"-"`
		IsError  bool        `json:"error"`
		Data     interface{} `json:"data"`
		Meta     interface{} `json:"meta"`
	}

	HTTPResponse interface {
		StatusCode() int
	}
)

func (h HTTPResponseWrapper) StatusCode() int {
	return h.HttpCode
}

func (h HTTPResponseWrapper) ToJson() ([]byte, error) {
	r, err := json.Marshal(h)
	return r, err
}

func JsonResponse(w http.ResponseWriter, statusCode int, val interface{}, meta interface{}) {
	w.Header().Set("content-type", ContentTypeJson)
	w.WriteHeader(statusCode)

	var succes = false
	if statusCode == http.StatusOK || statusCode == http.StatusTeapot {
		succes = true
	}

	json.NewEncoder(w).Encode(HTTPResponseWrapper{
		HttpCode: statusCode,
		IsError:  succes,
		Data:     val,
		Meta:     meta,
	})

}
