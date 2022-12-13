package models

import "time"

type Session struct {
	Username string `json:"username"`
	Value string `json:"value"`
	Expiry time.Time `json:"expiry"`
}

func (s *Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}