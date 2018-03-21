package controller

import (
	"encoding/json"
	"net/http"
)

type index struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func (i index) registerHandlers() {
	http.HandleFunc("/", i.indexHandler)
}

func (i index) indexHandler(w http.ResponseWriter, r *http.Request) {
	p := index{Status: "ok", Message: "breakfast is served"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&p)
}