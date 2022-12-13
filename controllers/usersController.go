package controllers

import (
	"net/http"
	"os"
	"path/filepath"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// read all users
	jsonData, err := os.ReadFile(filepath.Join("db", "users.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(string(jsonData)))
}