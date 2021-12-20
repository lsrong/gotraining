package handlers

import (
	"encoding/json"
	"net/http"
)

func Routers() {
	http.HandleFunc("/json", sendJson)
}

func sendJson(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name  string
		Email string
	}{
		"Ls",
		"Ls@gmail.com",
	}

	/*
		blog := struct {
			Title   string
			Content string
			Author  string
		}{
			"Go Blog",
			"Object Oriented Programming Mechanics",
			"demo",
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&user)
	//_ = user
	//_ = json.NewEncoder(w).Encode(&blog)
}
