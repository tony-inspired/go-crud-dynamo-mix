package http

import (
	"net/http"
	"encoding/json"
)

type response struct {
	Status int `json:"status"`
	Result interface{} `json:"result"`
}

//not a response level function
func newResponse(data interface{}, status int) *response {
	return &response{
		Status: status,
		Result: data,
	}
}

func (resp *response) bytes() []byte {
	data, _ := json.Marshal(resp)	
	return data	
}

func (resp *response)sendResponse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusTeapot)

	w.Write(resp.bytes())
}

func StatusOk(w http.ResponseWriter, r *http.Request, data interface{}){
	newResponse(data, http.StatusOK).sendResponse(w, r)
}