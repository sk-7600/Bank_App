package web

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w *http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statusCode, response)
}

func writeToHeader(w *http.ResponseWriter, statusCode int, payload interface{}) {
	(*w).WriteHeader(statusCode)
	(*w).Write(payload.([]byte))
}

func WriteErrorInResponse(err error, w http.ResponseWriter) {
	if err != nil {
		var x = []byte(err.Error())
		w.Write(x)
	}
}
