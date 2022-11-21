package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, code int, payload interface{}) {
	response, _ :=json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ResponseError(w http.ResponseWriter, code int, message string){
	ResponseSuccess(w, code, map[string]string{"message": message})
	
}