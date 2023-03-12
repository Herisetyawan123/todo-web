package handler

import (
	"encoding/json"
	"net/http"
)

type ResWelcome struct {
	Title string `json:"title"`
	Msg   string `json:"message"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(ResWelcome{
		Title: "Welcome to Aplication Programming Interface",
		Msg:   "I Hope you, enjoyed!",
	})
}
