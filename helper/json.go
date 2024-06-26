package helper

import (
	"encoding/json"
	"net/http"
)

type WebResponse struct {
	//Code   int         `json:"code"`
	Status string      `json:"message"`
	Data   interface{} `json:"data,omitempty"`
}

func ReadRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponse(write http.ResponseWriter, response WebResponse, status int) {
	write.Header().Add("Content-Type", "application/json")
	write.WriteHeader(status)
	if status != http.StatusNoContent {
		//write.WriteHeader(http.StatusText(response.Status))
		encoder := json.NewEncoder(write)
		err := encoder.Encode(response)
		PanicIfError(err)
	}
}
