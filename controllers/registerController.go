package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/bazoka-kaka/golang-todo-list/models"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// read all users
	var users []models.Credentials

	jsonData, err := os.ReadFile(filepath.Join("db", "users.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(jsonData, &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// append the users
	users = append(users, creds)

	// write the users
	jsonData, err = json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(filepath.Join("db", "users.json"), jsonData, 0666); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Register success!"))
}