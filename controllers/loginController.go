package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bazoka-kaka/golang-todo-list/models"
	"github.com/google/uuid"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
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

	// check creds
	userFound := false
	for _, user := range users {
		if user.Username == creds.Username && user.Password == creds.Password {
			userFound = true
			break
		}
	}

	if !userFound {
		http.Error(w, "Wrong username or password!", http.StatusUnauthorized)
		return
	}

	// read all sessions
	var sessions []models.Session

	jsonData, err = os.ReadFile(filepath.Join("db", "sessions.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(jsonData, &sessions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// create new session
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(time.Minute * 30)

	newSession := models.Session{
		Username: creds.Username,
		Value: sessionToken,
		Expiry: expiresAt,
	}
	
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: expiresAt,
	})

	// append new session
	sessions = append(sessions, newSession)

	// return sessions to json type
	jsonData, err = json.Marshal(sessions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write updated sessions
	if err := os.WriteFile(filepath.Join("db", "sessions.json"), jsonData, 0666); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Login Success!\n\nsession_token=%s\n\n", sessionToken)))
}