package server

import (
	//"fmt"
	"net/http"
	"encoding/json"
)

type Message struct{
	Info string `json:"version"`
}

// a simple get requests that return hello as string
func Info(w http.ResponseWriter, req *http.Request) {

	u := Message{"v:0.0.1"}
	uM, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(uM)

}
