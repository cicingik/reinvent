package rest

import (
	"net/http"

	"github.com/cicingik/reinvent/pkg/httpresp"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	httpresp.JsonResponse(w, http.StatusTeapot, nil, struct {
		Message string `json:"message"`
	}{
		Message: http.StatusText(http.StatusTeapot),
	})
}
