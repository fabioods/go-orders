package response

import (
	"encoding/json"
	"net/http"

	"github.com/fabioods/go-orders/pkg/errorformatted"
)

func WriteResponse(w http.ResponseWriter, data interface{}, err error, defaultStatus int) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		ef := err.(*errorformatted.ErrorFormatted)
		w.WriteHeader(ef.Status)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(defaultStatus)
		json.NewEncoder(w).Encode(data)
	}

}
