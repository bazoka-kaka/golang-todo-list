package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bazoka-kaka/golang-todo-list/models"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "You are not logged in!", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad request!", http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// read all sessions
	var sessions []models.Session
	
	jsonData, err := os.ReadFile(filepath.Join("db", "sessions.json"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(jsonData, &sessions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// find user session
	var userSession models.Session
	newSessions := []models.Session{}

	for _, session := range sessions {
		if session.Value == sessionToken {
			userSession = session
			continue
		}
		newSessions = append(newSessions, session)
	}

	nilSession := models.Session{}
	if userSession == nilSession {
		http.Error(w, "You are not logged in!", http.StatusUnauthorized)	
		return
	}

	// write new session
	jsonData, err = json.Marshal(newSessions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := os.WriteFile(filepath.Join("db", "sessions.json"), jsonData, 0666); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Path: "/",
		Value: "",
		Expires: time.Now(),
	})

	w.Write([]byte("Logout success!"))
}